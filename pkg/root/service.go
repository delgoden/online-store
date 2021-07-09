package root

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

// GiveRoleAdministrator gives the user the admin role
func (s *Service) GiveRoleAdministrator(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}

// RemoveRoleAdministrator removes the administrator role from the user
func (s *Service) RemoveRoleAdministrator(ctx context.Context, id int64) (*types.Status, error) {
	return nil, nil
}
