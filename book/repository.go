package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(ID int) error
	Authenticated(user LoginCredentials) (LoginCredentials, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(ID int) (Book, error) {
	var book Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Create(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book Book) (Book, error) {

	err := r.db.Model(&book).Updates(book).Error

	return book, err
}

func (r *repository) Delete(ID int) error {

	err := r.db.Delete(Book{}, ID).Error

	return err
}

func (r *repository) Authenticated(user LoginCredentials) (LoginCredentials, error) {
	var authuser LoginCredentials

	r.db.Where(&LoginCredentials{Email: user.Email, Password: user.Password}, "Email", "Password").Find(&authuser)

	return authuser, nil

}
