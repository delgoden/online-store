package product

import (
	"context"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

func (s *Service) GetCategories(ctx context.Context) ([]types.Category, error) {
	return nil, nil
}

func (s *Service) GetProducts(ctx context.Context) ([]types.Product, error) {
	return nil, nil
}

func (s *Service) GetProductsInCategory(ctx context.Context, categoryID int64) ([]types.Product, error) {
	return nil, nil
}

func (s *Service) GetProductByID(ctx context.Context, id int64) (*types.Product, error) {
	return nil, nil
}
