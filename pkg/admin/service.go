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
	return category, nil
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	return product, nil
}

// UpdateProduct updates existing products
func (s *Service) UpdateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	return nil, nil
}

// RemoveProduct removes product
func (s *Service) RemoveProduct(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}

// AddFoto adds a new photo
func (s *Service) AddFoto(ctx context.Context, foto *types.Foto) (*types.Status, error) {
	return nil, nil
}

// RemoveFoto deletes photo
func (s *Service) RemoveFoto(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}
