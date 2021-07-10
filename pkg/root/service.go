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
	status := &types.Status{}
	_, err := s.pool.Exec(ctx, `UPDATE users SET role = 'ADMINISTRATOR' WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	status.Status = true
	return status, nil
}

// RemoveRoleAdministrator removes the administrator role from the user
func (s *Service) RemoveRoleAdministrator(ctx context.Context, id int64) (*types.Status, error) {
	status := &types.Status{}
	_, err := s.pool.Exec(ctx, `UPDATE users SET role = 'CUSTOMER' WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	status.Status = true
	return status, nil
}
