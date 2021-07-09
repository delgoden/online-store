package root

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

func (s *Service) GiveRoleAdministrator(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}

func (s *Service) RemoveRoleAdministrator(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}
