package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func fileUploader(res http.ResponseWriter, req *http.Request) {
	file, header, _ := req.FormFile("fileuploader") // Here we call the FormFile handler on the
	// HTTP request to get the file for the provided form key.
	defer file.Close() // The defer statement closes the file once we return from the function.

	out, _ := os.Create(header.Filename) // Here we are creating a file
	//  inside the same directory with mode 666, which means the client can read
	// and write but cannot execute the file.
	defer out.Close()
	io.Copy(out, file) //  Here we copy content from the file we received to the file we created inside the same directory.
	fmt.Fprintf(res, "file uploaded successfully "+header.Filename)

}
func index(res http.ResponseWriter, req *http.Request) {
	parseTemp, _ := template.ParseFiles("templates/index.html")
	parseTemp.Execute(res, nil)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/upload", fileUploader).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}
