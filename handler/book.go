package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (bh *bookHandler) GetAllBooksHandler(c *gin.Context) {
	var bks []book.Book

	bks, err := bh.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range bks {

		br := bookMapping(b)

		booksResponse = append(booksResponse, br)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})

}

func (bh *bookHandler) FindBookById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	bk, err := bh.bookService.FindById(i)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	bookResponse := bookMapping(bk)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (bh *bookHandler) AddBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		} else {
			errorMessage := fmt.Sprintln("Another error")
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	bk, err := bh.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := bookMapping(bk)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (bh *bookHandler) UpdateBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		} else {
			errorMessage := fmt.Sprintln("Another error")
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	bk, err := bh.bookService.Update(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := bookMapping(bk)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (bh *bookHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	err := bh.bookService.Delete(i)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Successfully Deleted",
	})

}

func bookMapping(bk book.Book) book.BookResponse {

	bookResponse := book.BookResponse{
		ID:          bk.ID,
		Title:       bk.Title,
		Author:      bk.Author,
		Price:       bk.Price,
		Description: bk.Description,
		Rating:      bk.Rating,
	}

	return bookResponse
}

/*
   func (bh *bookHandler) RootHandler(c *gin.Context) {
   	c.JSON(http.StatusOK, gin.H{
   		"name": "Praditio Aditya Nugraha",
   		"bio":  "Software Quality Engineer",
   	})
   }

   func (bh *bookHandler) HelloHandler(c *gin.Context) {
   	c.JSON(http.StatusOK, gin.H{
   		"title":   "Hello World!",
   		"content": "This is to test golang web api.",
   	})

   }

   func (bh *bookHandler) BooksHandler(c *gin.Context) {
   	id := c.Param("id")
   	title := c.Param("title")

   	c.JSON(http.StatusOK, gin.H{
   		"id":    id,
   		"title": title,
   	})
   }

   func (bh *bookHandler) QueryHandler(c *gin.Context) {
   	id := c.Query("title")
   	price := c.Query("price")

   	c.JSON(http.StatusOK, gin.H{
   		"title": id,
   		"price": price,
   	})
   }*/
