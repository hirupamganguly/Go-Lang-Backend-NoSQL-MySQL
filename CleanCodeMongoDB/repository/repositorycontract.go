package repository

import (
	"../entity"
)

// BookRepositoryContract ...
type BookRepositoryContract interface {
	Find(name string) (*entity.Book, error)
	Store(*entity.Book) error
}
