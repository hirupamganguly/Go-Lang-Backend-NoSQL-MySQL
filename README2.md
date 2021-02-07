
### REST-API-CRUD
```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo" // Package mongo provides a MongoDB Driver API for Go.
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	//Unlike a language like Java (where there are only primitive and reference types),
	//Go has types to represent textual, numeric, boolean, pointer, composite, function,
	// and interface values. Once a variable is declared to be of a certain type, 
	//  it can only carry values of that type.
	//Package primitive contains types similar to Go primitives for BSON types
	//   can do not have direct Go primitive representations.
)

// -----------DATA MODEL---------------

// Author is...
type Author struct {
	FName string `json:"fname,omitempty" bson:"fname,omitempty"`
	LName string `json:"lname,omitempty" bson:"lname,omitempty"`
}

// Book is...
type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Author *Author            `json:"author,omitempty" bson:"author,omitempty"`
}

// ------x-----DATA MODEL-------x--------

//------------HANDELER FUNCTIONS----------------
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Content types also known as MIME type or media types are a two part identifier for file formats.
	// The HTTP header Content-Type is responsible for telling the HTTP client or server what type of data is being sent.
	// To specify the content types of the request body and output, use the Content-Type and Accept headers. 
	//Indicates that the request body format is JSON.
	var books []Book
	collection := client.Database("booksmanagement").Collection("books")
	cursor, _ := collection.Find(ctx, bson.M{})
	// The Cursor is a MongoDB Collection of the document which is returned upon the find method execution.
	//By default, it is automatically executed as a loop. However, 
	//  we can explicitly get specific index document from being returned cursor.
	//It is just like a pointer which is pointing upon a specific index value.
	// In simple words when we call a find method, all the documents which are returned are saved in a virtual cursor.

	defer cursor.Close(ctx) // Instructs the server to close a cursor and free associated server resources.
	for cursor.Next(ctx) {  // The cursor.Next() method is used to return the next document in a cursor.
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
	fmt.Println("showed")
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	params := mux.Vars(r)  // Vars returns the route variables for the current request, if any.
	id, _ := primitive.ObjectIDFromHex(params["id"]) 
	// ObjectIDFromHex creates a new ObjectID from a hex string.
	//   It returns an error if the hex string is not a valid ObjectID.
	collection := client.Database("booksmanagement").Collection("books")
	collection.FindOne(ctx, Book{ID: id}).Decode(&book)
	json.NewEncoder(w).Encode(book)

}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	collection := client.Database("booksmanagement").Collection("books")
	result, _ := collection.InsertOne(ctx, book)
	json.NewEncoder(w).Encode(result)
	fmt.Println("added")
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	json.NewDecoder(r.Body).Decode(&book)
	collection := client.Database("booksmanagement").Collection("books")
	upd := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"fname", book.Author.FName},
				{"lname", book.Author.LName},
			}}}},
	}
	//D is an ordered representation of a BSON document. This type should be used when the
	// order of the elements matters, such as MongoDB command documents.
	// If the order of the elements does not matter, an M should be used instead.
	//E represents a BSON element for a D. It is usually used inside a D
	//M is an unordered representation of a BSON document. This type should be used when the
	//order of the elements does not matter. This type is handled as a regular map[string]
	//interface{} when encoding and decoding. Elements will be serialized in an undefined, random order.
	//ObjectID is the BSON ObjectID type.
	//NewObjectID generates a new ObjectID.
	//ObjectIDFromHex creates a new ObjectID from a hex string. It returns an error if the
	//hex string is not a valid ObjectID.

	collection.FindOneAndUpdate(ctx, Book{ID: id}, upd).Decode(&book)
	json.NewEncoder(w).Encode(book)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("booksmanagement").Collection("books")
	res, _ := collection.DeleteOne(ctx, Book{ID: id})
	json.NewEncoder(w).Encode(res)
}

//-----x------HANDELER FUNCTIONS-------x--------

var client *mongo.Client 
// Client is a handle representing a pool of connections to a MongoDB deployment. 
//  It is safe for concurrent use by multiple goroutines.
// The Client type opens and closes connections automatically and maintains a pool of idle connections.
// For connection pool configuration options, see documentation for the ClientOptions type in the mongo/options package.
var err error
var ctx, _ = context.WithTimeout(context.Background(), 1560*time.Second)

func main() {

	// create routes:-
	router := mux.NewRouter()
	// ---------API/BOOKS ROUTES-------------
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/createbook", createBook).Methods("POST")
	router.HandleFunc("/api/updatebook/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/deletebook/{id}", deleteBook).Methods("DELETE")
	// -----X----API/BOOKS ROUTES-------X------

	fmt.Println("Starting...")
	// -------mongo-atlas connection-----------
	// Client is a handle representing a pool of connections to a MongoDB deployment.
	// It is safe for concurrent use by multiple goroutines.
	// NewClient creates a new client to connect to a deployment specified by the uri.  This includes the ApplyURI method.
	// The Client type opens and closes connections automatically and maintains a pool of idle connections.
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rupam:<password>@cluster0.cpwla.mongodb.net/<db-name>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx) 
	// Connect initializes the Client by starting background monitoring goroutines. 
	//If the Client was created using the NewClient function, this method must be called before a Client can be
	// used.Connect starts background goroutines to monitor the state of the deployment and
	//  does not do any I/O in the main goroutine.
	 //The Client.Ping method can be used to verify that the connection was created successfully.
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	 // Disconnect closes sockets to the topology referenced by this Client.
	// It will shut down any monitoring goroutines, close the idle connection pool, and will wait until all the in use
	// connections have been returned to the connection pool and closed before returning.
	// If the context expires via cancellation, 
	//deadline, or timeout before the in use connections have returned, 
	//  the in use connections will be closed, resulting in the failure
	// of any in flight read or write operations. If this method returns with no errors,
	// all connections associated with this Client have been closed.
	// ----x---mongo-atlas connection----x-------
	// -------------GO-LOCALHOST SERVER CONNECTION----------------------
	http.ListenAndServe(":12345", router)
	// -------X-----GO-LOCALHOST SERVER CONNECTION---------X------------
	fmt.Println("established.... :)")
}

```
<img src="ASSETS/postman-create-book.PNG">
<img src="ASSETS/mongodb-atlas-create-book.PNG">


### MySQL-REST-API-CRUD

```go
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
	password  = "password"
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
	db, _ = sql.Open("mysql", "uic42w7p2k7lydc5:password@tcp(bnd5jogxwkdk3abcibgx-mysql.services.clever-cloud.com:3306)/bnd5jogxwkdk3abcibgx")
}
func deleteRecord(res http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	name, _ := vals["name"]
	stmt, _ := db.Prepare("DELETE from employee where name=?")
	result, _ := stmt.Exec(name[0])// DELETE from employee where name=name[0]
	rowsAffected, _ := result.RowsAffected()
	fmt.Fprintf(res, "Rows Affected %d", rowsAffected)

}
func updateRecord(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	vals := req.URL.Query()
	name, _ := vals["name"]
	smt, _ := db.Prepare("UPDATE employee SET name=? where uid=?")
	result, _ := smt.Exec(name[0], id) // UPDATE employee SET name=name[0] where uid=id
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
	result, _ := stmt.Exec(name[0])// INSERT employee SET name=name[0]
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
```

<img src="ASSETS/MysQlCRUDdemoVideo.gif">

## Clean-code-architecture

Sometime we really need to meet requirements and hit the deadlines. But even in this case we should keep our project at least well structured, readable and extendable.
An architecture should not depend on frameworks, you should be able to swap a framework with the least effort.
A software system’s core logic should not be affected by changes in UI, Databases, Frameworks, Libraries, etc.
An inner layer should not know anything about upper/outer layers. As a result dependencies can only point inwards. In particular, the name of something declared in an outer circle must not be mentioned by the code in an inner circle. That includes, functions, classes. variables, or any other named software entity.

You should always pay attention to objects being passed between layers. An object being passed should be isolated, simple or even just a plain data type without hidden dependencies. 