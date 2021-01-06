// Go’s net/http package offers a lot of functionalities for URL routing of the HTTP
// requests. One thing it doesn’t do very well is dynamic URL routing. Fortunately, we
// can achieve this with the gorilla/mux package

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// First, we defined GetRequestHandler and PostRequestHandler, which simply write a
// message on an HTTP response stream

//GetRequestHandler ...
var GetRequestHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("get request handler"))
})

//PostRequestHandler ...
var PostRequestHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("post request handler"))
})

// we defined PathVariableHandler, which extracts request path variables, gets the
// value, and writes it to an HTTP response stream

//GetRequestPATHVarHandler ...
var GetRequestPATHVarHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	response.Write([]byte(" Get Request PATH VarHandler -> " + name))
})

func main() {
	// 	Once we run the program, the HTTP server will start locally listening on port 8080, and
	// accessing http://localhost:8080/, http://localhost:8080/poster, and
	// http://localhost:8080/hello/Rupam Ganguly from a browser or command line will produce the
	// message defined in the corresponding handler definition.
	router := mux.NewRouter()
	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/poster", PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}", GetRequestPATHVarHandler).Methods("GET")
	http.ListenAndServe("localhost:8080", router)
}
