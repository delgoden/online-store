package user

import (
	"context"
	"errors"
	"log"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrNotFound          = errors.New("item not found")
	ErrNotEnoughProducts = errors.New("not enough products")
)

type Service struct {
	pool *pgxpool.Pool
}

// NewService constructor function to create the service
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

func (s *Service) Buy(ctx context.Context, position *types.Position, userID int64) (*types.Purchase, error) {
	qty := 0

	purchases := &types.Purchase{
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
		if err == pgx.ErrNoRows {
			log.Println(err)
			return nil, ErrNotFound
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}
		_, err = s.pool.Exec(ctx, `UPDATE products SET qty = qty - $1 WHERE id =$2`, position.Qty, position.ProductID)
		if err != nil {
			return nil, err
		}

		purchases.PositionsID = position.ID
		log.Println(purchases)
	}

	err = s.pool.QueryRow(ctx,
		`INSERT INTO purchases (user_id, position_id)
		VALUES ($1, $2)
		RETURNING id, created`, purchases.UserID, purchases.PositionsID).Scan(&purchases.ID, &purchases.Created)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return nil, ErrNotFound
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return purchases, nil
}

func (s *Service) AddProductIntoCart(ctx context.Context, position *types.Position, userID int64) (*types.Status, error) {
	qty := 0
	status := &types.Status{}
	err := s.pool.QueryRow(ctx, `SELECT qty FROM products WHERE id = $1`, position.ProductID).Scan(&qty)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if qty > position.Qty {
		_, err = s.pool.Exec(ctx, `
			INSERT INTO cart (user_id, product_id, qty, price) 
			VALUES ($1, $2, $3, $4)`, userID, position.ProductID, position.Qty, position.Price)
		if err != nil {
			return nil, err
		}
		status.Status = true
	}
	return status, nil
}

func (s *Service) RemoveProductFromCart(ctx context.Context, productID, userID int64) (*types.Status, error) {
	_, err := s.pool.Exec(ctx, `DELETE FROM cart WHERE product_id = $1 AND user_id =$2`, productID, userID)
	if err != nil {
		return nil, err
	}
	status := &types.Status{
		Status: true,
	}
	return status, nil
}

func (s *Service) BuyFromCart(ctx context.Context, userID int64) ([]types.Purchase, error) {
	rows, err := s.pool.Query(ctx, `SELECT product_id, qty, price FROM cart WHERE user_id =$1`, userID)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return nil, ErrNotFound
	}

	defer rows.Close()
	purchases := []types.Purchase{}
	for rows.Next() {
		position := &types.Position{}

		err := rows.Scan(&position.ProductID, &position.Qty, &position.Price)
		if err == pgx.ErrNoRows {
			log.Println(err)
			return nil, ErrNotFound
		}
		if err != nil {
			return nil, err
		}
		qty := 0
		err = s.pool.QueryRow(ctx, `SELECT qty FROM products WHERE id = $1`, position.ProductID).Scan(&qty)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if qty > position.Qty {

			err = s.pool.QueryRow(ctx,
				`INSERT INTO positions (product_id, qty, price)
			VALUES ($1, $2, $3)
			RETURNING id`, position.ProductID, position.Qty, position.Price).Scan(&position.ID)
			if err == pgx.ErrNoRows {
				log.Println(err)
				return nil, ErrNotFound
			}
			if err != nil {
				return nil, err
			}

			_, err = s.pool.Exec(ctx, `UPDATE products SET qty = qty - $1 WHERE id =$2`, position.Qty, position.ProductID)
			if err != nil {
				return nil, err
			}

			purchase := &types.Purchase{
				UserID:      userID,
				PositionsID: position.ID,
			}
			err = s.pool.QueryRow(ctx,
				`INSERT INTO purchases (user_id, position_id)
			VALUES ($1, $2)
			RETURNing id, created`, userID, position.ID).Scan(&purchase.ID, &purchase.Created)
			if err == pgx.ErrNoRows {
				log.Println(err)
				return nil, ErrNotFound
			}
			if err != nil {
				return nil, err
			}
			purchases = append(purchases, *purchase)

			_, err = s.pool.Exec(ctx, `DELETE FROM cart WHERE user_id =$1`, userID)
			if err != nil {
				return nil, err
			}
		} else {
			return purchases, ErrNotEnoughProducts
		}
	}
	return purchases, nil
}
