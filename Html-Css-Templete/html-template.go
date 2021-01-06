package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

//  Here we define a person struct type that
// has Id and Name fields.

// Person ...
type Person struct {
	Id   string
	Name string
}

// RenderTemplate ...
var RenderTemplate = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	fmt.Println("hi")
	person := Person{Id: "1134", Name: "Rupam Ganguly"}
	// Here we are calling ParseFiles of the html/template package, which creates a new template and
	// parses the filename we pass as an input, which is first-template.html ,
	// The resulting template will have the name and contents of the input file.
	// Html-Css-Templete\HTMLS\template.html
	parsTemp, _ := template.ParseFiles("templates/template.html")
	parsTemp.Execute(response, person)

	// 	 parsedTemplate.Execute(w, person): Here we are calling an Execute handler on a
	// parsed template, which injects person data into the template, generates an HTML
	// output, and writes it onto an HTTP response stream.
})

func main() {
	router := mux.NewRouter()
	router.Handle("/", RenderTemplate).Methods("GET")
	// 	PathPrefix adds a matcher for the URL path prefix. This matches if the given template is a prefix of the full URL path.
	// Note that it does not treat slashes specially ("/foobar/" will be matched by the prefix "/foo") so you may want to use a trailing slash here.
	//When you specify a path using PathPrefix() it has an implicit wildcard at the end.
	//On the other hand, when you specify a path using Path(), there's no such implied wildcard suffix.
	router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))
	http.ListenAndServe("localhost:8080", router)
}

// Assume that
// I have a file

// /home/go/src/js/kor.js
// Then, tell fileserve serves local directory

// fs := http.FileServer(http.Dir("/home/go/src/js"))
// Example 1 - root url to Filerserver root
// Now file server takes "/" request as "/home/go/src/js"+"/"

// http.Handle("/", fs)
// Yes, http://localhost/kor.js request tells Fileserver, find kor.js in

// "/home/go/src/js" +  "/"  + "kor.js".
// we got kor.js file.
