package videomicro

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// // The transport.go is responsible for converting a transport layer request into an endpoint call.

// CreateRequest ...
type CreateRequest struct {
	v VideoModel
}

// CreateResponse ..,
type CreateResponse struct {
	Msg string `json:"msg"`
	Err error
}

// GetRequest ...
type GetRequest struct {
}

// GetResponse ...
type GetResponse struct {
	Data interface{} `json:"video"`
	Err  error       `json:"error,omitempty"`
}

// UpdateRequest ...
type UpdateRequest struct {
	ID string `json:"id"`
	v  VideoModel
}

// UpdateResponse ...
type UpdateResponse struct {
	Msg string `json:"msg,omitempty"`
	Err error  `json:"error,omitempty"`
}

// DeleteRequest ...
type DeleteRequest struct {
	ID string `json:"id"`
}

// DeleteResponse ...
type DeleteResponse struct {
	Msg string `json:"msg,omitempty"`
	Err error  `json:"error,omitempty"`
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req.v); err != nil {
		return nil, err
	}
	return req, nil
}
func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetRequest
	return req, nil
}
func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateRequest
	vars := mux.Vars(r)
	req = UpdateRequest{
		ID: vars["id"],
	}
	if err := json.NewDecoder(r.Body).Decode(&req.v); err != nil {
		return nil, err
	}
	return req, nil
}
func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DeleteRequest
	vars := mux.Vars(r)
	req = DeleteRequest{
		ID: vars["id"],
	}
	return req, nil
}
func encodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(rw).Encode(response)
}
