package admin

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

func (s *Service) CreateCategory(ctx context.Context, category *types.Category) (*types.Category, error) {
	return nil, nil
}

func (s *Service) UpdateCategory(ctx context.Context, category *types.Category) (*types.Category, error) {
	return nil, nil
}

func (s *Service) CreateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	return nil, nil
}

func (s *Service) UpdateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	return nil, nil
}

func (s *Service) RemoveProduct(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}

func (s *Service) AddFoto(ctx context.Context, foto *types.Foto) (*types.Status, error) {
	return nil, nil
}
func (s *Service) RemoveFoto(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}
