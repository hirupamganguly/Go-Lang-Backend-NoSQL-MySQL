package usecases

import (
	"time"

	"../entity"
	"../repository"
)

type bookService struct {
	bookRepo repository.BookRepositoryContract
}

// NewBookService ...
func NewBookService(bookRepo repository.BookRepositoryContract) BookServicesContract {
	return &bookService{bookRepo}
}

func (bs *bookService) Find(name string) (*entity.Book, error) {
	return bs.bookRepo.Find(name)
}

func (bs *bookService) Store(book *entity.Book) error {
	book.ID = time.Now().UTC().Unix()
	return bs.bookRepo.Store(book)
}
