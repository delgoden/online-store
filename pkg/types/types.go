package types

import "time"

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
	Create   time.Time `json:"create"`
}

// Category is a structure defining a product category
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Product represents the structure defining the product
type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CategoryID int64     `json:"category_id"`
	FotoID     []int64   `json:"foto_id"`
	Qty        int       `json:"qty"`
	Price      int       `json:"price"`
	Create     time.Time `json:"create"`
	Update     time.Time `json:"update"`
	Active     bool      `json:"active"`
}

// Foto is a structure for transferring a photo
type Foto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
