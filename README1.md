
## Binary-search-tree

```go
package main

import "fmt"

// Node ...
type Node struct {
	key  int
	data string
	l    *Node
	r    *Node
}

// Bst ...
type Bst struct {
	root *Node
}

// derefferncing will done by the compiler automatically
func (n *Node) insertRecursively(node Node) {
	// function to find the correct place for a node in a tree

	// when new node is lesser than n -> go to left subtree of n and
	//again
	//start the function by making - left child of n ->as n
	// restore Point(Previous Execution State) stored at callstack.. it restored again when
	// newly function call finshed..
	// when new node is greater than n -> go to right subtree of n and
	//again
	//start the function by making - right child of n ->as n
	// restore Point(Previous Execution State) stored at callstack.. it restored again when
	// newly function call finshed..
	if n == nil {
		return //  base condition of Recursive Problem->    if(n==null) return ;
	} else if node.key <= n.key {
		if n.l == nil {
			n.l = &node
		} else {
			n.l.insertRecursively(node)
		}
	} else {
		if n.r == nil {
			n.r = &node
		} else {
			n.r.insertRecursively(node)
		}
	}
}

func (bst *Bst) insert(keyy int, value string) *Bst {
	// insert function create Node then add the node to its correct position by calling insertRecursively()

	node := &Node{key: keyy, data: value}
	if bst.root == nil {
		bst.root = &Node{key: keyy, data: value}
	} else {
		bst.root.insertRecursively(*node)
	}
	return bst
}
func searchRecursively(keyy int, n *Node) bool {

	if n == nil {
		return false //// base case of recursive problem.
	}
	// recursively call for left subtree and right subtree, by checking the key is smaller or greater than n
	//(as it is recursive function- here n variable contains actually curent node each time..)
	// when key is lesser than n -> go to left subtree of n and
	//again
	//start the function by making - left child of n ->as n
	// when key is greater than n -> go to right subtree of n and
	// again
	//start the function by making - right child of n ->as n
	if keyy < n.key {
		return searchRecursively(keyy, n.l)
	}
	if keyy > n.key {
		return searchRecursively(keyy, n.r)
	}
	// we find the match so return true
	return true
}

func (bst *Bst) search(keyy int) bool {
	return searchRecursively(keyy, bst.root)
}
func removeRecursively(keyy int, n *Node) *Node {
	if n == nil {
		// base case of recursive Problem.
		return nil
	}
	if keyy < n.key {
		// find the element at left subtree
		n.l = removeRecursively(keyy, n.l)
		return n
	}
	if keyy > n.key {
		// find the element at right subtree
		n.r = removeRecursively(keyy, n.r)
		return n
	}
	//when found then:-
	if keyy == n.key {
		// no child
		if n.l == nil && n.r == nil {
			n = nil
			return nil
		}
		// single child
		if n.l == nil {
			// if it has only right child then replaced it with right child.
			n = n.r
			return n
		}
		if n.r == nil {
			// if it has only left child then replaced it with left child.
			n = n.l
			return n
		}
		// two children

		// try to find the next small node after the node which we want to delete
		// here we create and initialize lmostOfRight variable as -> n.right, means we just have to find
		// the smallest of n.right, because that will be the in-Order Successor
		lmostOfRight := n.r
		for {
			if lmostOfRight != nil && lmostOfRight.l != nil {
				lmostOfRight = lmostOfRight.l // just incrementing towards left...
			} else {
				break
			}
		}
		// Found it,
		//then... replace with it
		//as it has both child then replaced it with in order successor. then delete that inorder successor as
		// it is still prsent at its position.
		//Now the twist is as it is inorder successor so it can have right child or no chils,
		// as it is left most of n.right.
		// Now again we call delete function by passing the inorder successor's key.
		n.key, n.data = lmostOfRight.key, lmostOfRight.data
		//then... delete it
		n.r = removeRecursively(n.key, n.r)
		return n
	}
	return n
}
func (bst *Bst) remove(keyy int) {
	removeRecursively(keyy, bst.root)
}

func (bst *Bst) smallest() {
	n := bst.root
	for {
		if n.l != nil {
			n = n.l
		} else {
			fmt.Println("Smallest item is: ", n.data, n.key)
			break
		}
	}
}
func (bst *Bst) largest() {
	n := bst.root
	for {
		if n.r != nil {
			n = n.r
		} else {
			fmt.Println("largest item is: ", n.data, n.key)
			break
		}
	}

}

type list struct {
	items []Node
}

func (lst *list) push(item Node) {
	lst.items = append(lst.items, item)
}

func (lst *list) popQueue() *Node {
	if len(lst.items) == 0 {
		return nil
	}
	temp := lst.items[0]
	lst.items = lst.items[1:]
	return &temp
}

func (bst *Bst) bfs() {
	fmt.Print("BFS ")
	if bst.root == nil {
		return
	}
	var sq list
	sq.push(*bst.root)
	for {
		if len(sq.items) == 0 {
			fmt.Println("")
			return
		}

		n := sq.popQueue()
		fmt.Print(" -> ", n.key, " ", n.data)
		if n.l != nil {
			sq.push(*n.l)
		}
		if n.r != nil {
			sq.push(*n.r)
		}
	}

}

func (lst *list) popStack() *Node {
	if len(lst.items) == 0 {
		return nil
	}
	temp := lst.items[len(lst.items)-1]
	lst.items = lst.items[:len(lst.items)-1]
	return &temp
}

func (bst *Bst) dfs() {
	fmt.Print("DFS ")
	if bst.root == nil {
		return
	}
	var sq list
	sq.push(*bst.root)
	for {
		if len(sq.items) == 0 {
			fmt.Println("")
			return
		}

		n := sq.popStack()
		fmt.Print(" -> ", n.key, " ", n.data)
		if n.l != nil {
			sq.push(*n.l)
		}
		if n.r != nil {
			sq.push(*n.r)
		}
	}

}

func main() {
	var b Bst
	b.insert(12, "hi")
	b.insert(2, "hi")
	b.insert(21, "hi")
	b.insert(32, "hi")
	b.insert(42, "hi")
	b.insert(3, "hi")
	b.insert(27, "hi")
	b.insert(107, "hi")
	b.insert(31, "hi")
	b.largest()
	b.smallest()
	b.bfs()
	b.dfs()
	b.remove(32)
	fmt.Println("delete 32")
	b.bfs()
	b.remove(12)
	fmt.Println("delete 12")
	b.bfs()
	b.dfs()
}

```
```shell
largest item is:  hi 107
Smallest item is:  hi 2
BFS  -> 12 hi -> 2 hi -> 21 hi -> 3 hi -> 32 hi -> 27 hi -> 42 hi -> 31 hi -> 107 hi
DFS  -> 12 hi -> 21 hi -> 32 hi -> 42 hi -> 107 hi -> 27 hi -> 31 hi -> 2 hi -> 3 hi
delete 32
BFS  -> 12 hi -> 2 hi -> 21 hi -> 3 hi -> 42 hi -> 27 hi -> 107 hi -> 31 hi
delete 12
BFS  -> 21 hi -> 2 hi -> 42 hi -> 3 hi -> 27 hi -> 107 hi -> 31 hi
DFS  -> 21 hi -> 42 hi -> 107 hi -> 27 hi -> 31 hi -> 2 hi -> 3 hi

```

## Backend-Development

In this whole Note I will try to avoid error handeling as it is so easy that you just replace '_' to err
then you have to add this if block :

```go
if err != nil
{
 log.Printf("error is : ", err)
 return
}

```
This will print the error .

### Gorilla-Mux-demo
```go
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
```
### HTML-CSS-TEMPLATE-RENDERING
HTML -

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Template</title>
    <link href="styles/style.css" type="text/css" rel="stylesheet"/>
</head>
<body>
    <!-- The preceding template has two placeholders, {{.Name}} and {{.Id}},
	
	 whose
values will be substituted or injected by the template engine at runtime. -->

    <h1>Hello {{.Name}} </h1>
    Your ID is {{.Id}}
</body>
</html>
```
CSS -

```css
body{
    color: blue;
}
```

Go -

```go
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
	// Here we are calling ParseFiles of the templates/template.html package, which creates a new template and
	// parses the filename we pass as an input, which is template.html ,
	// The resulting template will have the name and contents of the input file.
	
	parsTemp, _ := template.ParseFiles("templates/template.html")
	parsTemp.Execute(response, person)

	// 	 parsedTemplate.Execute(response, person): Here we are calling an Execute handler on a
	// parsed template, which injects person data into the template, generates an HTML
	// output, and writes it onto an HTTP response stream.
})

func main() {
	router := mux.NewRouter()
	router.Handle("/", RenderTemplate).Methods("GET")
	// 	PathPrefix adds a matcher for the URL path prefix.
	// This matches if the given template is a prefix of the full URL path.
	// Note that it does not treat slashes specially ("/foobar/" will be matched by the 
	//  prefix "/foo") so you may want to use a trailing slash here.
	//When you specify a path using PathPrefix() it has an implicit wildcard at the end.
	//On the other hand, when you specify a path using Path(), there's no such implied wildcard suffix.
	router.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))
	// actually this is little complicated...
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
```


<img src="ASSETS\html-template-rendering.PNG">

### Interact-with-Form

<img src="ASSETS/form-structures.PNG">

form.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link href="styles/main.css" type="text/css" rel="stylesheet"/>
</head>
<body>
    <h1>Log In</h1>
    <form method="POST" action="/login">
        <label for="ussername">Name</label>
        <input type="text" id="uname" name="username">
        <label for="password">Password</label>
        <input type="password" id="pword" name="password">
        <label for="email">email</label>
        <input type="email" id="eml" name="email">
        <p>type Comment : </p>
        <textarea name="comment"placeholder="Remember, be nice!" cols="30" rows="5"></textarea>
        <p></p>
        <button type="submit">LOGIN</button>
    </form>
</body>
</html>
```

homepage.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>are Logged In {{.Username}}</h1> 
    <p>
        your comment is {{.Comment}}</p>
</body>
</html>
```
main.css

```css
label{
    color: darkcyan;
}
button{
    background-color: darkgreen;
    color: floralwhite;
    width: 200px;
    height: 3em;
    text-align: center;
    display: inline;
    font-size: medium;
}
textarea {
    width: 400px;
    height: 20em;
	}
```

form.go

```go
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
```

### File-Upload

index.html

```html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <form action="/upload" method="POST" enctype="multipart/form-data">
        <input type="file" name="fileuploader">
        <input type="submit" name="submit" value="Submit">
    </form>

</body>
</html>

```
fileupload.go

```go
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
	io.Copy(out, file) //  Here we copy content from the file we received to the file
	//   we created inside the same directory.
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

```

