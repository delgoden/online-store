package auth

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

// SignUp user registration
func (s *Service) SignUp(ctx context.Context, signUpData *types.Auth) (*types.User, error) {
	return nil, nil
}

// SignIn user authorization
func (s *Service) SignIn(ctx context.Context, signInData *types.Auth) (*types.Token, error) {
	return nil, nil
}
