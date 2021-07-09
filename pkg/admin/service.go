package admin

import (
	"context"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx/v4/pgxpool"
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
	return nil, nil
}

// UpdateCategory updates an existing category
func (s *Service) UpdateCategory(ctx context.Context, category *types.Category) (*types.Category, error) {
	return nil, nil
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	return nil, nil
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
