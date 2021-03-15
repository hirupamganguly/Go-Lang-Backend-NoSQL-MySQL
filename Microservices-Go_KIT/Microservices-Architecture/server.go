package mymicroservice

import (
	"context"
	"net/http"

	ht "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("GET").Path("/status").Handler(ht.NewServer(endpoints.StatusEndpoint, decodeStatusReq, encodeRes))
	r.Methods("GET").Path("/get").Handler(ht.NewServer(endpoints.GetDataEndpoint, decodeGetDataReq, encodeRes))
	return r
}

// Handlers are responsible for writing response headers and bodies. Almost any object can be a handler, so long as it satisfies the http.Handler interface. In lay terms, that simply means it must have a ServeHTTP method with the following signature:
// ServeHTTP(http.ResponseWriter, *http.Request)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)

		// Go's HTTP server takes in an address to listen on, and a handler. Internally, it creates a TCP listener to accept connections on the given address, and whenever a request comes in, it:

		// Parses the raw HTTP request (path, headers, etc) into a http.Request
		// Creates a http.ResponseWriter for sending the response
		// Invokes the handler by calling its ServeHTTP method, passing in the Request and ResponseWriter
	})
}

// Custom Handlers
// Let's create a custom handler which responds with the current local time in a given format:

// type timeHandler struct {
//   format string
// }

// func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//   tm := time.Now().Format(th.format)
//   w.Write([]byte("The time is: " + tm))
// }
// All that really matters is that we have an object (in this case it's a timeHandler struct, but it could equally be a string or function or anything else), and we've implemented a method with the signature ServeHTTP(http.ResponseWriter, *http.Request) on it. That's all we need to make a handler.
