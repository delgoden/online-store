package types

import (
	"mime/multipart"
	"time"
)

// Auth is a structure for transferring registration and authorization data
type Auth struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Status to transfer the status of execution or not execution of the request
type Status struct {
	Status bool `json:"status"`
}

// Token to transfer the token
type Token struct {
	Token string `json:"token"`
}

// User is a user-defining structure
type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
}

type UserToken struct {
	Token   string    `json:"token"`
	UserID  int64     `json:"user_id"`
	Created time.Time `json:"created"`
	Expire  time.Time `json:"expire"`
}

// Category is a structure defining a product category
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Product represents the structure defining the product
type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CategoryID  int64     `json:"category_id"`
	Description string    `json:"description"`
	PhotosURL   []string  `json:"photos_url"`
	Qty         int       `json:"qty"`
	Price       int       `json:"price"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Active      bool      `json:"active"`
}

// Foto is a structure for transferring a photo
type Photo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	File multipart.File
}

type Position struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	Qty       int   `json:"qty"`
	Price     int   `json:"price"`
}

type Purchase struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	PositionsID int64     `json:"position_id"`
	Created     time.Time `json:"created"`
}

type Cart struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Qty       int   `json:"qty"`
	Price     int   `json:"price"`
}
