package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Employee ...
type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Employees ...
type Employees []Employee

var employees []Employee
var employeesV1 []Employee
var employeesV2 []Employee

func init() {
	employees = Employees{Employee{Id: "1", FirstName: "Rupam", LastName: "Ganguly"}, Employee{Id: "2", FirstName: "Rintu", LastName: "Ganguly"}}

	employeesV1 = Employees{Employee{Id: "3", FirstName: "RupamV1", LastName: "GangulyV1"}, Employee{Id: "4", FirstName: "RintuV1", LastName: "GangulyV1"}}

	employeesV2 = Employees{Employee{Id: "5", FirstName: "RupamV2", LastName: "GangulyV2"}, Employee{Id: "6", FirstName: "RintuV2", LastName: "GangulyV2"}}
}

// Route ...
type Route struct {
	Name        string
	Method      string
	pattern     string
	HandlerFunc http.HandlerFunc
}

func getEmployees(res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/v1") { // localhost:8080/v1/employees
		json.NewEncoder(res).Encode(employeesV1)
	} else if strings.HasPrefix(req.URL.Path, "/v2") { // localhost:8080/v2/employees
		json.NewEncoder(res).Encode(employeesV2)
	} else { // localhost:8080/employees
		json.NewEncoder(res).Encode(employees)
	}
}
func getEmployee(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	for _, empl := range employees {
		if empl.Id == id {
			json.NewEncoder(res).Encode(empl)
		}
	}
}

//Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"getEmployees", "GET", "/employees", getEmployees,
	},
	Route{
		"getEmployee", "GET", "/employee/{id}", getEmployee,
	},
}

// we defined an AddRoutes helper function, which iterates over the routes array
// we defined, adds it to the gorilla/mux router, and returns the Router object.

// AddRoutes ...
func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).Path(route.pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	// V1
	AddRoutes(muxRouter.PathPrefix("/v1").Subrouter())

	// V2
	AddRoutes(muxRouter.PathPrefix("/v2").Subrouter())
	http.ListenAndServe("localhost:8080", router)
}
