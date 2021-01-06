package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
)

const (
	//USER ...
	USER = "admin"
	//PASSWORD ...
	PASSWORD = "passw"
)

func helloWeb(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello Web")
}

// BAuth is...
func BAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {

	return func(response http.ResponseWriter, request *http.Request) {
		// 		we first get the username and password provided in
		// the request's authorization header using r.BasicAuth() then compare it to the
		// constants declared in the program.
		// If credentials match, then it returns the
		// handler, otherwise it sets WWW-Authenticate along with a status code of 401 and
		// writes You are Unauthorized to access the application on an HTTP response stream.
		user, pass, ok := request.BasicAuth()
		fmt.Println(user, pass, ok)
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(PASSWORD)) != 1 {
			response.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			response.WriteHeader(401)
			response.Write([]byte("You are Unauthorized to access the application.\n"))
			// The HTTP WWW-Authenticate response header defines the authentication method that should be used to
			// gain access to a resource. The WWW-Authenticate header is sent along with a 401 Unauthorized response.
			// Authentication type. A common type is "Basic". IANA maintains a list of Authentication schemes.
			// realm=A description of the protected area. If no realm is specified, clients often display a formatted hostname instead.
			return
		}
		handler(response, request)
	}
}
func main() {
	http.HandleFunc("/", BAuth(helloWeb, "Please enter your username and password"))
	http.ListenAndServe("localhost:8080", nil)
}
