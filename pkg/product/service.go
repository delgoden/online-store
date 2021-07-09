package product

import (
	"context"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	pool *pgxpool.Pool
}

// NewService constructor function to create the service
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

// GetCategories gives a list of existing categories
func (s *Service) GetCategories(ctx context.Context) ([]types.Category, error) {
	return nil, nil
}

// GetProducts displays a complete list of products
func (s *Service) GetProducts(ctx context.Context) ([]types.Product, error) {
	return nil, nil
}

// GetProductsInCategory displays a list of products in a category
func (s *Service) GetProductsInCategory(ctx context.Context, categoryID int64) ([]types.Product, error) {
	return nil, nil
}

// GetProductByID issues the product according to its ID
func (s *Service) GetProductByID(ctx context.Context, id int64) (*types.Product, error) {
	return nil, nil
}
