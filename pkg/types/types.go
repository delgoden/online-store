package types

import "time"

type Auth struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Status struct {
	Status bool `json:"status"`
}

type Token struct {
	Token string `json:"token"`
}

type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Create   time.Time `json:"create"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

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

type Foto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
