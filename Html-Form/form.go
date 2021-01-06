package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//User ...
type User struct { // DATA MODEL
	Username string
	password string
	email    string
	Comment  string
}

func formValidatorStringCount(user *User) bool {
	if len(user.Username) < 3 || len(user.password) < 6 {
		return false
	}
	return true
}
func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isWord(s string) bool {
	for _, c := range s {
		if !isLetter(c) {
			return false
		}
	}
	return true
}
func isGmail(s string) bool {
	if strings.ContainsAny(s, "@gmail.com") {
		return true
	}
	return false
}
func readForm(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user := new(User) // user object created
	fmt.Println("username - ", req.FormValue("username"))
	fmt.Println("username - ", req.FormValue("comment"))
	user.Username = req.FormValue("username")
	user.password = req.FormValue("password")
	user.Comment = req.FormValue("comment")
	user.email = req.FormValue("email")

	if !formValidatorStringCount(user) {
		fmt.Fprint(res, `<script type="text/javascript"  charset="utf-8">
		alert("You have to enter at least 6 characters for Password and 3 letters for Name!");
		</script>`)
		return
	}

	if !isWord(user.Username) {
		fmt.Fprint(res, `<script type="text/javascript"  charset="utf-8">
		alert("Expecting only letters in Name!");
		</script>`)
		return
	}
	if isGmail(user.email) {
		fmt.Fprint(res, `<script type="text/javascript"  charset="utf-8">
		alert("Expecting only valid gmail in Email!");
		</script>`)
		return
	}
	parseTemp, _ := template.ParseFiles("templates/htmls/homePage.html")
	parseTemp.Execute(res, user)
}

func login(res http.ResponseWriter, req *http.Request) {
	parseTemp, _ := template.ParseFiles("templates/htmls/form.html")
	parseTemp.Execute(res, nil)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", login).Methods("GET")
	router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))
	router.HandleFunc("/login", readForm).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}
