package user

import (
	"context"
	"log"

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

func (s *Service) Buy(ctx context.Context, position *types.Position, userID int64) (*types.Purchases, error) {
	qty := 0

	purchases := &types.Purchases{
		UserID: userID,
	}
	err := s.pool.QueryRow(ctx, `SELECT qty FROM products WHERE id = $1`, position.ProductID).Scan(&qty)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if qty > position.Qty {

		err = s.pool.QueryRow(ctx,
			`INSERT INTO
			positions (product_id, qty, price)
			VALUES ($1, $2, $3)
			RETURNING id`, position.ProductID, position.Qty, position.Price).Scan(&position.ID)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		purchases.PositionsID = position.ID
	}

	err = s.pool.QueryRow(ctx,
		`INSERT INTO purchases (user_id, position_id)
		VALUES ($1, $2)
		RETURNING id, created`, purchases.UserID, purchases.PositionsID).Scan(purchases.ID, purchases.Created)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return purchases, nil
}
