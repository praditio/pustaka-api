package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:12345678@tcp(localhost:3300)/pustaka-api-perf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	//db.AutoMigrate(&book.Book{})
	//db.AutoMigrate(&book.LoginCredentials{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	jwtService := book.JWTAuthService()
	loginService := book.NewLoginService(bookRepository)
	loginHandler := handler.LoginHandler(loginService, jwtService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/login", loginHandler.Login)
	v1.POST("/books/add", middleware.AuthorizeJWT(), bookHandler.AddBooksHandler)
	v1.PUT("/books/update", middleware.AuthorizeJWT(), bookHandler.UpdateBooksHandler)
	v1.GET("/books/all", middleware.AuthorizeJWT(), bookHandler.GetAllBooksHandler)
	v1.GET("/books/:id", middleware.AuthorizeJWT(), bookHandler.FindBookById)
	v1.DELETE("/books/:id", middleware.AuthorizeJWT(), bookHandler.Delete)

	router.Run(":8080")
}

/*bookRequest := book.BookRequest{}
bookRequest.Title = "Filosofi Teras"
bookRequest.Author = "Henry Manampiring"
bookRequest.Rating = 5
bookRequest.Price = "75000"
bookRequest.Description = "Bercerita tentang pemahaman stoisisme."

bookService.Create(bookRequest)*/

//CREATE
/*	db.AutoMigrate(&model.Book{})

	book := model.Book{}
	book.Title = "Atomic Habits"
	book.Price = 100000
	book.Author = "Agus Setiawan"
	book.Rating = 4
	book.Description = "Self development about building good habits."

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("Error creating book record")
	}*/

//READ
/*var books []model.Book

err = db.Debug().Where("Title ='Man Tiger'").Find(&books).Error
if err != nil {
	fmt.Println("Failed retrieving book record")
}

for _, b := range books {
	fmt.Println("Title: ", b.Title)
	fmt.Println("Author: ", b.Author)
}*/

//UPDATE
/*var book model.Book

err = db.First(&book, 1).Error
if err != nil {
	fmt.Println("Not found")
}

book.Title = "Man Tiger (Revised Edition)"
err = db.Save(&book).Error
if err != nil {
	fmt.Println("Not successfully updated")
}*/

//DELETE
/*var book model.Book

err = db.First(&book, 1).Error
if err != nil {
	fmt.Println("Not found")
}

err = db.Delete(&book).Error
if err != nil {
	fmt.Println("Not successfully deleted")
}*/

//v1.GET("/", bookHandler.RootHandler)
//v1.GET("/hello", bookHandler.HelloHandler)
//v1.GET("/books/:id", bookHandler.FindBookById)
//v1.GET("/query", bookHandler.QueryHandler)
