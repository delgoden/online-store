package admin

import (
	"context"
	"errors"
	"log"

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
func (s *Service) UpdateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	return nil, nil
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

// AddFoto adds a new photo
func (s *Service) AddFoto(ctx context.Context, foto *types.Foto) (*types.Status, error) {
	return nil, nil
}

// RemoveFoto deletes photo
func (s *Service) RemoveFoto(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}
