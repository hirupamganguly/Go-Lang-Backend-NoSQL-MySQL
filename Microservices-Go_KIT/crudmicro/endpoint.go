package crudmicro

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	CreateUserEndpoint  endpoint.Endpoint
	GetUserByIDEndpoint endpoint.Endpoint
	GetAllUsersEndpoint endpoint.Endpoint
	DeleteUserEndpoint  endpoint.Endpoint
	UpdateUserEndpoint  endpoint.Endpoint
}

// // Endpoints have a simple definition. They take in a context and a request, and return a response and an error.

// MakeCreateUserEndpoint ...
func MakeCreateUserEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req.user)
		return CreateUserResponse{ID: id, Err: err}, nil
	}

}

// MakeGetUserByIDEndpoint ...
func MakeGetUserByIDEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIDRequest)
		fmt.Println("Request", req)
		id := req.ID
		data, err := s.GetUserByID(ctx, id)
		fmt.Println("ID decoded output:", data)
		return GetUserByIDResponse{Data: data, Err: err}, nil
	}

}

// MakeGetAllUsersEndpoint ...
func MakeGetAllUsersEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		data, err := s.GetAllUsers(ctx)
		return GetAllUsersResponse{Data: data, Err: err}, nil
	}

}

// MakeDeleteUserEndpoint ...
func MakeDeleteUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		fmt.Println("Request of DeleteUser", req)
		fmt.Println("Rquest ud:", req.ID)
		id := req.ID
		msg, err := s.DeleteUser(ctx, id)
		return DeleteUserResponse{Msg: msg, Err: err}, nil
	}

}

// MakeUpdateUserEndpoint ...
func MakeUpdateUserEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		fmt.Println("Request", req.user)
		rc := req.user
		fmt.Println("Request Id", rc.ID)
		fmt.Println("REQ.user:", req.user)
		err := s.UpdateUser(ctx, rc.ID.String(), req.user)
		return UpdateUserResponse{Err: err}, nil
	}
}
