package handler

import "net/http"

// BookHandler ...
type BookHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}
