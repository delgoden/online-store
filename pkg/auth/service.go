package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"

	"github.com/delgoden/internet-store/pkg/types"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInternal        = errors.New("internal error")
	ErrNoSuchUser      = errors.New("no such user")
	ErrLoginUsed       = errors.New("login already registred")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokernNotFound  = errors.New("token not found")
	ErrTokenExpired    = errors.New("token expired")
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
	hash, err := bcrypt.GenerateFromPassword([]byte(signUpData.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternal
	}
	user := &types.User{}
	err = s.pool.QueryRow(ctx, `
		INSERT INTO users (name, login, password)
		VALUES ($1, $2, $3)		
		ON CONFLICT DO NOTHING RETURNING id, name, login, role, created
	`, signUpData.Name, signUpData.Login, hash).Scan(&user.ID, &user.Name, &user.Login, &user.Role, &user.Created)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return nil, ErrLoginUsed
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

// SignIn user authorization
func (s *Service) SignIn(ctx context.Context, signInData *types.Auth) (*types.Token, error) {
	var (
		hash string
		id   int64
	)
	token := &types.Token{}
	err := s.pool.QueryRow(ctx, `SELECT id, password FROM users WHERE login = $1`, signInData.Login).Scan(&id, &hash)
	if err == pgx.ErrNoRows {
		log.Println(err)
		return token, ErrNoSuchUser
	}
	if err != nil {
		log.Println(err)
		return token, ErrInternal
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(signInData.Password))
	if err != nil {
		log.Println(err)
		return token, ErrInvalidPassword
	}

	buffer := make([]byte, 256)
	n, err := rand.Read(buffer)
	if n != len(buffer) || err != nil {
		log.Println(err)
		return token, ErrInternal
	}

	token.Token = hex.EncodeToString(buffer)
	_, err = s.pool.Exec(ctx, `INSERT INTO users_tokens(token, user_id) VALUES ($1, $2)`, token.Token, id)
	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}

	return token, nil
}
