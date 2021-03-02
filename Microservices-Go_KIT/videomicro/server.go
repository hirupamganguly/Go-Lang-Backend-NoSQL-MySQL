package videomicro

import (
	"context"
	"net/http"

	ht "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// // Handlers are responsible for writing response headers and bodies. Almost any object can be a handler, so long as it satisfies the http.Handler interface. In lay terms, that simply means it must have a ServeHTTP method with the following signature:
// ServeHTTP(http.ResponseWriter, *http.Request)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/v/create").Handler(ht.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeResponse))

	r.Methods("GET").Path("/v/get").Handler(ht.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeResponse))

	r.Methods("PUT").Path("/v/update/{id}").Handler(ht.NewServer(endpoints.UpdateEndpoint, decodeUpdateRequest, encodeResponse))

	r.Methods("DELETE").Path("/v/delete/{id}").Handler(ht.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeResponse))

	return r
}
