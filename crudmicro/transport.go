package crudmicro

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// The transport.go is responsible for converting a transport layer request into an endpoint call.

type (
	// CreateUserRequest ...
	CreateUserRequest struct {
		user User
	}
	// CreateUserResponse ...
	CreateUserResponse struct {
		ID  string `json:"id"`
		Err error
	}
	// GetUserByIDRequest ...
	GetUserByIDRequest struct {
		ID string `json:"id"`
	}
	// GetUserByIDResponse ...
	GetUserByIDResponse struct {
		Data interface{} `json:"user"`
		Err  error       `json:"error,omitempty"`
	}
	// GetAllUsersRequest ...
	GetAllUsersRequest struct {
	}
	// GetAllUsersResponse ...
	GetAllUsersResponse struct {
		Data interface{} `json:"user"`
		Err  error       `json:"error,omitempty"`
	}
	// DeleteUserRequest ...
	DeleteUserRequest struct {
		ID string `json:"id"`
	}
	// DeleteUserResponse ...
	DeleteUserResponse struct {
		Msg string `json:"msg,omitempty"`
		Err error  `json:"error,omitempty"`
	}
	// UpdateUserRequest ...
	UpdateUserRequest struct {
		ID   string `json:"id"`
		user User
	}
	// UpdateUserResponse ...
	UpdateUserResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.user); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetUserByIDRequest
	vars := mux.Vars(r)
	req = GetUserByIDRequest{
		ID: vars["id"],
	}

	return req, nil
}

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllUsersRequest

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DeleteUserRequest
	vars := mux.Vars(r)
	req = DeleteUserRequest{
		ID: vars["id"],
	}

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	vars := mux.Vars(r)
	req = UpdateUserRequest{
		ID: vars["id"],
	}
	if err := json.NewDecoder(r.Body).Decode(&req.user); err != nil {
		return nil, err
	}
	return req, nil

}

//  encodes the output
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
