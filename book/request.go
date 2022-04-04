package book

import "encoding/json"

type BookRequest struct {
	ID          int         `json:"id"`
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json:"rating" binding:"number"`
	Author      string      `json:"author" binding:"required"`
}

type LoginCredentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
