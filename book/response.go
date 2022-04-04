package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

type LoginCredentialsResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
