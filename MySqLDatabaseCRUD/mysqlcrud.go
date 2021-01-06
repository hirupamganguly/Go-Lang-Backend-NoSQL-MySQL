package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Employee ...
type Employee struct {
	Id   int    `json:"uid"`
	Name string `json:"name"`
}

// I am using cloud remote mysql server from clever cloud
const (
	username  = "uic42w7p2k7lydc5"
	password  = "RWHma741Oz6JszOmF6TW"
	hostname  = "bnd5jogxwkdk3abcibgx-mysql.services.clever-cloud.com"
	port      = "3306"
	dbname    = "bnd5jogxwkdk3abcibgx"
	conString = "" + username + ":" + password + "@tcp(" + hostname + ":" + port + ")" + "/" + dbname
	//"user7:s$cret@tcp(127.0.0.1:3306)/testdb"
	//uic42w7p2k7lydc5:password@tcp(bnd5jogxwkdk3abcibgx-mysql.services.clever-cloud.com:3306)/bnd5jogxwkdk3abcibgx
)

var version string
var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "uic42w7p2k7lydc5:RWHma741Oz6JszOmF6TW@tcp(bnd5jogxwkdk3abcibgx-mysql.services.clever-cloud.com:3306)/bnd5jogxwkdk3abcibgx")
}
func deleteRecord(res http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	name, _ := vals["name"]
	stmt, _ := db.Prepare("DELETE from employee where name=?")
	result, _ := stmt.Exec(name[0])
	rowsAffected, _ := result.RowsAffected()
	fmt.Fprintf(res, "Rows Affected %d", rowsAffected)

}
func updateRecord(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	vals := req.URL.Query()
	name, _ := vals["name"]
	smt, _ := db.Prepare("UPDATE employee SET name=? where uid=?")
	result, _ := smt.Exec(name[0], id) // "UPDATE employee SET name[0] where uid=id"
	rowsAffected, _ := result.RowsAffected()
	fmt.Fprintf(res, "No of Rows Affected %d ", rowsAffected)
}
func readRecords(res http.ResponseWriter, req *http.Request) {
	rows, _ := db.Query("SELECT * FROM employee")
	employees := []Employee{}
	for rows.Next() {
		var uid int
		var name string
		rows.Scan(&uid, &name)
		emp := Employee{Id: uid, Name: name}
		employees = append(employees, emp)

	}
	json.NewEncoder(res).Encode(employees)
}
func createRecord(res http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	name, _ := vals["name"]
	stmt, _ := db.Prepare("INSERT employee SET name=?")
	result, _ := stmt.Exec(name[0])
	id, _ := result.LastInsertId()
	fmt.Fprintf(res, " LAST INSERTED RECORD HAVE ID:: %s", strconv.FormatInt(id, 10))
}
func getCurrentDb(res http.ResponseWriter, req *http.Request) {
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Fprintf(res, version)
}
func getCurrentConnString(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, conString)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getCurrentConnString).Methods("GET")
	router.HandleFunc("/db", getCurrentDb).Methods("GET")
	router.HandleFunc("/emp/create", createRecord).Methods("POST")
	router.HandleFunc("/emp/read", readRecords).Methods("GET")
	router.HandleFunc("/emp/update/{id}", updateRecord).Methods("PUT")
	router.HandleFunc("/emp/delete", deleteRecord).Methods("DELETE")
	defer db.Close()
	http.ListenAndServe("localhost:8080", router)
}
