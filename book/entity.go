package book

import "time"

type Book struct {
	ID          int
	Title       string
	Author      string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LoginCredentials struct {
	Email    string
	Password string
}
