package test

import (
	"pustaka-api/book"
	"pustaka-api/book/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBookService_GetBookById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBookRepo := mocks.NewMockRepository(mockCtrl)
	mockBookRepo.EXPECT().FindById(8).Return(book.Book{
		ID:          8,
		Title:       "Perahu Kertas",
		Author:      "Dewi Lestari",
		Description: "Perjalanan Kugy mencari agen neptunus.",
		Price:       90000,
		Rating:      5,
	}, nil)

	bookExpectedResponse := book.Book{
		ID:          8,
		Title:       "Perahu Kertas",
		Author:      "Dewi Lestari",
		Description: "Perjalanan Kugy mencari agen neptunus.",
		Price:       90000,
		Rating:      5,
	}

	bookService := book.NewService(mockBookRepo)
	bookActualResponse, _ := bookService.FindById(bookExpectedResponse.ID)

	if bookActualResponse != bookExpectedResponse {
		t.Errorf("The data returned does not match.")
	}

}

func TestBookService_GetTotalPrice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBookRepo := mocks.NewMockRepository(mockCtrl)
	mockBookRepo.EXPECT().FindAll().Return([]book.Book{{
		ID:          8,
		Title:       "Perahu Kertas",
		Author:      "Dewi Lestari",
		Description: "Perjalanan Kugy mencari agen neptunus.",
		Price:       90000,
		Rating:      5,
	}, {
		ID:          2,
		Title:       "Atomic Habits",
		Author:      "Agus Setiawan",
		Description: "Self development about building good habits.",
		Price:       100000,
		Rating:      4,
	}, {
		ID:          3,
		Title:       "Laskar Pelangi",
		Author:      "Andrea Hirata",
		Description: "Buku ini mengisahkan 6 sekawan yang tinggal di Bangka Belitung",
		Price:       150000,
		Rating:      5,
	},
	}, nil)

	bookService := book.NewService(mockBookRepo)
	bookActualResponse, _ := bookService.GetTotalPrice()
	bookExpectedResponse := 340000

	if bookActualResponse != bookExpectedResponse {
		t.Errorf("Wrong Calculation!")
	}

}
