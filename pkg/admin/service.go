package admin

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"sync"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrInternal              = errors.New("internal error")
	ErrCategoryAlreadyExists = errors.New("category already exists")
	ErrCategoryDoesNotExist  = errors.New("category does not exist")
	ErrProductAlreadyExists  = errors.New("product already exists")
	ErrProductDoesNotExist   = errors.New("product does not exist")
)

// Service
type Service struct {
	pool *pgxpool.Pool
}

// NewService constructor function to create the service
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

// CreateCategory  creates a new category
func (s *Service) CreateCategory(ctx context.Context, category *types.Category) (*types.Category, error) {
	err := s.pool.QueryRow(ctx, `
	INSERT INTO categories (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, category.Name).Scan(&category.ID)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return nil, ErrCategoryAlreadyExists
	}
	if err != nil {
		log.Println(err)
		return category, err
	}
	return category, nil
}

// UpdateCategory updates an existing category
func (s *Service) UpdateCategory(ctx context.Context, category *types.Category) (*types.Category, error) {
	test := &types.Category{}
	err := s.pool.QueryRow(ctx, `SELECT name FROM categories WHERE id = $1`, category.ID).Scan(&test.Name)
	if err == pgx.ErrNoRows || test.Name == "" {
		//log.Println(err, category)
		category.ID = 0
		return category, ErrCategoryDoesNotExist
	}

	_, err = s.pool.Exec(ctx, `UPDATE categories SET name = $1 WHERE id = $2`, category.Name, category.ID)
	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return category, nil
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	err := s.pool.QueryRow(ctx, `
            INSERT INTO products (name, category_id, description, qty, price)
            VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT DO NOTHING 
            RETURNING id, created, updated, active
        `, product.Name, product.CategoryID, product.Description, product.Qty, product.Price).
		Scan(&product.ID, &product.Created, &product.Updated, &product.Active)
	if err == pgx.ErrNoRows && product.ID == 0 {
		log.Println(err)
		return nil, ErrProductAlreadyExists
	}
	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return product, nil
}

// UpdateProduct updates existing products
func (s *Service) UpdateProduct(ctx context.Context, product *types.Product) (*types.Status, error) {
	active := false
	name := ""
	status := &types.Status{
		Status: false,
	}
	err := s.pool.QueryRow(ctx, `SELECT name, active FROM products WHERE id = $1`, product.ID).Scan(&name, &active)
	if err == pgx.ErrNoRows && name == "" {
		log.Println(err, status.Status)
		return status, ErrProductDoesNotExist
	}

	query := " UPDATE products SET"
	if product.Name != "" {
		query += fmt.Sprintf(" name = '%v'", product.Name)
	}
	if product.CategoryID != 0 {
		query += fmt.Sprintf(", category_id = %v", product.CategoryID)
	}
	if product.Description != "" {
		query += fmt.Sprintf(", description = '%v'", product.Description)
	}
	if product.Qty != 0 {
		query += fmt.Sprintf(", qty = %v", product.Qty)
	}
	if product.Price != 0 {
		query += fmt.Sprintf(", price = '%v'", product.Price)
	}
	query += ", updated = CURRENT_TIMESTAMP"
	if product.Active != active {

		query += fmt.Sprintf(", active = %t", product.Active)
	}
	query += " WHERE id = $1"
	log.Println(query)
	_, err = s.pool.Exec(ctx, query, product.ID)
	if err != nil {
		return status, err
	}

	status.Status = true
	return status, nil
}

// RemoveProduct removes product
func (s *Service) RemoveProduct(ctx context.Context, id int64) (*types.Status, error) {
	product := &types.Product{}
	status := &types.Status{}
	err := s.pool.QueryRow(ctx, `SELECT name FROM products WHERE id = $1`, id).Scan(&product.Name)
	if product.Name == "" {
		log.Println(err)
		return status, ErrProductDoesNotExist
	}

	_, err = s.pool.Exec(ctx, `
            DELETE FROM products WHERE id = $1
        `, id)
	if err != nil {
		log.Println(err)
		return status, ErrInternal
	}
	status.Status = true
	return status, nil
}

// Addphoto adds a new photo
func (s *Service) AddPhoto(ctx context.Context, photo *types.Photo, productID int64) (*types.Status, error) {

	ch := make(chan error)
	status := &types.Status{}
	name := photo.Name
	err := s.pool.QueryRow(ctx, `SELECT name FROM products WHERE id =$1`, productID).Scan(&photo.Name)
	if err == pgx.ErrNoRows {
		return status, ErrProductDoesNotExist
	}
	photo.Name += name

	go func() {

		err = saveFile(photo.File, photo.Name)
		if err != nil {
			log.Print(err)
			ch <- err
			return
		}
	}()

	err = <-ch
	close(ch)
	if err != nil {
		return status, err
	}
	err = s.pool.QueryRow(ctx, `INSERT INTO photos (name, product_id) VALUES ($1, $2)`, photo.Name, productID).Scan(&photo.ID)
	if err == pgx.ErrNoRows {
		return status, err
	}
	if err != nil {
		return status, err
	}

	status.Status = true
	return status, nil
}

// Removephoto deletes photo
func (s *Service) RemovePhoto(ctx context.Context, photoID int64) (*types.Status, error) {
	wg := &sync.WaitGroup{}
	photo := &types.Photo{}
	status := &types.Status{
		Status: false,
	}
	err := s.pool.QueryRow(ctx, `DELETE FROM photos WHERE id = $1 RETURNING name`, photoID).Scan(&photo.Name)
	if err != nil {
		log.Println(err)
		return status, ErrInternal
	}

	wg.Add(1)
	go remoteBannerImage(wg, photo.Name)
	wg.Wait()
	status.Status = true
	return status, nil
}

func saveFile(data multipart.File, name string) error {
	f, err := os.Create("db/images" + name)
	if err != nil {
		log.Println("Can't open file: " + "db/images" + name)
		return err
	}
	defer f.Close()
	buf := make([]byte, 1024)

	for {
		n, err := data.Read(buf)

		if err != nil && err != io.EOF {
			log.Println("Couldn't write file: " + "db/images" + name)
			break
		}

		if n == 0 {
			break
		}

		if _, err := f.Write(buf[:n]); err != nil {
			log.Println("Couldn't write file: " + "db/images" + name)
			break
		}
	}
	return nil
}

func remoteBannerImage(wg *sync.WaitGroup, imageName string) {
	defer wg.Done()

	os.Remove("db/images" + imageName)
}
