package delivery

import (
	"encoding/json"

	"../entity"
)

type Book struct{}

func (b *Book) Decode(input []byte) (*entity.Book, error) {
	book := &entity.Book{}
	if err := json.Unmarshal(input, book); err != nil {
		return nil, err
	}
	return book, nil
}
func (b *Book) Encode(input *entity.Book) ([]byte, error) {
	msz, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return msz, nil
}
