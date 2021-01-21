package usecases

import (
	"../entity"
)

// BookServicesContract ...
type BookServicesContract interface {
	Find(name string) (*entity.Book, error)
	Store(*entity.Book) error
}
