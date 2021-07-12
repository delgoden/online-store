package product

import (
	"context"
	"errors"
	"log"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrNotFound            = errors.New("item not found")
	ErrInternal            = errors.New("internal error")
	ErrProductDoesNotExist = errors.New("product does not exist")
)

type Service struct {
	pool *pgxpool.Pool
}

// NewService constructor function to create the service
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

// GetCategories gives a list of existing categories
func (s *Service) GetCategories(ctx context.Context) ([]*types.Category, error) {
	categories := []*types.Category{}
	rows, err := s.pool.Query(ctx, `SELECT id name FROM categories`)
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}

	defer rows.Close()

	for rows.Next() {
		category := &types.Category{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println(err)
		}

		categories = append(categories, category)
	}
	err = rows.Err()
	if err != nil {
		return categories, ErrInternal
	}

	return categories, nil
}

// GetProducts displays a complete list of products
func (s *Service) GetAllActiveProducts(ctx context.Context) ([]*types.Product, error) {
	products := []*types.Product{}
	rowsProducts, err := s.pool.Query(ctx,
		`SELECT id, name, category_id, description, photo_id, qty, price FROM products WHERE active = true`)
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}

	defer rowsProducts.Close()

	for rowsProducts.Next() {
		product := &types.Product{}
		if err := rowsProducts.Scan(
			&product.ID, &product.Name, &product.CategoryID, &product.Description,
			&product.PhotosID, &product.Qty, &product.Price,
		); err != nil {
			log.Println(err)
		}

		rowsPhotos, err := s.pool.Query(ctx, `SELECT name FROM product_id = $1`, product.ID)
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		defer rowsPhotos.Close()

		for rowsPhotos.Next() {
			url := "http://localhost:9999/images/"
			photo := &types.Photo{}
			if err := rowsPhotos.Scan(&photo.Name); err != nil {
				log.Println(err)
			}
			product.PhotosURL = append(product.PhotosURL, url+photo.Name)
		}

		err = rowsPhotos.Err()
		if err != nil {
			return products, ErrInternal
		}

		products = append(products, product)
	}
	err = rowsProducts.Err()
	if err != nil {
		return products, ErrInternal
	}

	return products, nil
}

// GetProductsInCategory displays a list of products in a category
func (s *Service) GetProductsInCategory(ctx context.Context, categoryID int64) ([]*types.Product, error) {
	products := []*types.Product{}
	rowsProducts, err := s.pool.Query(ctx,
		`SELECT id, name, category_id, description, photo_id, qty, price FROM products WHERE active = true`)
	if err == pgx.ErrNoRows {
		return nil, ErrProductDoesNotExist
	}

	defer rowsProducts.Close()

	for rowsProducts.Next() {
		product := &types.Product{}
		if err := rowsProducts.Scan(
			&product.ID, &product.Name, &product.CategoryID, &product.Description,
			&product.PhotosID, &product.Qty, &product.Price,
		); err != nil {
			log.Println(err)
			return nil, ErrInternal
		}
		products = append(products, product)
	}
	err = rowsProducts.Err()
	if err != nil {
		return products, ErrInternal
	}

	return products, nil
}

// GetProductByID issues the product according to its ID
func (s *Service) GetProductByID(ctx context.Context, id int64) (*types.Product, error) {
	product := &types.Product{}
	err := s.pool.QueryRow(ctx,
		`SELECT id, name, category_id, description, photo_id, qty, price FROM products WHERE category_id = $1`, id).
		Scan(&product.ID, &product.Name, &product.CategoryID, &product.Description,
			&product.PhotosID, &product.Qty, &product.Price)
	if err != nil {
		return product, ErrProductDoesNotExist
	}
	return product, nil
}
