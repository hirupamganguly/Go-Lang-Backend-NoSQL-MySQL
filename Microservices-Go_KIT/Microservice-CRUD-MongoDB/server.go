package crudmicro

import (
	"context"
	"net/http"

	ht "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	//r.Use(commonMiddleware)

	r.Methods("POST").Path("/create").Handler(ht.NewServer(endpoints.CreateUserEndpoint, decodeCreateUserRequest, encodeResponse))

	r.Methods("GET").Path("/user/{id}").Handler(ht.NewServer(endpoints.GetUserByIDEndpoint, decodeGetUserByIDRequest, encodeResponse))

	r.Methods("GET").Path("/getall").Handler(ht.NewServer(endpoints.GetAllUsersEndpoint, decodeGetAllUsersRequest, encodeResponse))

	r.Methods("DELETE").Path("/deleteuser/{id}").Handler(ht.NewServer(endpoints.DeleteUserEndpoint, decodeDeleteUserRequest, encodeResponse))

	r.Methods("PUT").Path("/updateuser/{id}").Handler(ht.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, encodeResponse))
	return r
}

// Handlers are responsible for writing response headers and bodies. Almost any object can be a handler, so long as it satisfies the http.Handler interface. In lay terms, that simply means it must have a ServeHTTP method with the following signature:
// ServeHTTP(http.ResponseWriter, *http.Request)
