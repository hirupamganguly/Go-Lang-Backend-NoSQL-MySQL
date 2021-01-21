package handler

import (
	"log"
	"net/http"

	"../usecases"
)

type handler struct {
	bookService usecases.BookServicesContract
}

func NewHandler(bookService usecases.BookServicesContract) BookHandler {
	return &handler{bookService: bookService}
}
func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &jsondelivery.Redirect{}
}
