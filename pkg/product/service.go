package product

import (
	"context"
	"errors"
	"log"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrNotFound = errors.New("item not found")
	ErrInternal = errors.New("internal error")
)

type Service struct {
	pool *pgxpool.Pool
}

// NewService constructor function to create the service
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

// GetCategories gives a list of existing categories
func (s *Service) GetCategories(ctx context.Context) ([]*types.Category, error) {
	categories := []*types.Category{}
	rows, err := s.pool.Query(ctx, `SELECT id name FROM categories`)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	defer rows.Close()

	for rows.Next() {
		category := &types.Category{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println(err)
		}

		categories = append(categories, category)
	}
	err = rows.Err()
	if err != nil {
		return categories, ErrInternal
	}

	return categories, nil
}

// GetProducts displays a complete list of products
func (s *Service) GetAllProducts(ctx context.Context) ([]types.Product, error) {
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
