package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(bookRequest BookRequest) (Book, error)
	Delete(ID int) error
	GetTotalPrice() (int, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	book := bookMapping(bookRequest)

	return s.repository.Create(book)
}

func (s *service) Update(bookRequest BookRequest) (Book, error) {

	book := bookMapping(bookRequest)

	return s.repository.Update(book)
}

func (s *service) Delete(ID int) error {

	return s.repository.Delete(ID)
}

func (s *service) GetTotalPrice() (int, error) {

	allBooks, _ := s.repository.FindAll()
	sum := 0

	for _, b := range allBooks {
		sum += b.Price
	}

	return sum, nil
}

func bookMapping(br BookRequest) Book {

	price, _ := br.Price.Int64()

	book := Book{
		ID:          br.ID,
		Title:       br.Title,
		Price:       int(price),
		Description: br.Description,
		Author:      br.Author,
		Rating:      br.Rating,
	}

	return book

}
