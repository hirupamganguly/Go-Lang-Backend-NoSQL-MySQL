package delivery

import (
	"../entity"
)

type BookDelivery interface {
	Decode(input []byte) (*entity.Book, error)
	Encode(input *entity.Book) ([]byte, error)
}
y 