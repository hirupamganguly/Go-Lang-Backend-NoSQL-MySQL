# efficent-with-go-by-hirupam

### [ READ EVERY COMMENTS, THEN YOU WILL SURELY UNDERSTAND THE CONCEPTS ]
### [ I TRY TO MAKE IT WITH A SINGLE APPROACH, IS: AS FAST AS POSSIBLE TO LEARN GOLANG AND BACKEND DEVELOPMENT FROM A SINGLE RESOURCE ]

## Why-go-lang

Backend development languages handle the ‘behind-the-scenes’ functionality of an applications. It’s code that connects the application to a database, manages user connections, and powers the application itself. Backend developers typically earn higher salaries than front-end developers, as backend languages tend to be more technical.

Python: the most-loved programming language and The First suggestion for backend development in 2021.
Python is capable of unleashing its full potential when used in web development, machine learning, and fintech. Moreover, thanks to its scalability, it also proves to be popular among startups. You can check a whole range of Python’s applications in our extensive guide on what's it used for. So as you can see, regardless of whether you’re looking to build a large or small digital product, chances are, you’ll find a relevant Python framework and a knowledgeable team pretty fast.

The second suggestion for backend development in 2021 is a statically typed, compiling programming language: Go, also known as `Golang`. Released in 2009, it has been named the Programming Language of the Year by TIOBE twice so far and – while it’s surely less popular than the above-mentioned Python – it’s been mentioned as the fifth most loved technology by 62.3% of the Stack Overflow survey’s respondents in 2020. If you want to develop faster, perhaps because you have many different services to write, or you have a large team of developers, then Go is your language of choice.
Golang performs on average 40 times better than Python when making the same calculations. 
Python’s response time is longer. In the same timeframe (1 minute), we're able to handle even 10500 requests in Go, whereas in Python we arrive at as much as  2000 requests fewer. 

Rust and Go are both awesome
Go is fast, but Rust is faster.
Go has an efficient garbage collector, but Rust has static memory management.
Go has great concurrency support, but Rust has provably-correct concurrency.
Go has interfaces, but Rust has traits and other zero-cost abstractions.
You can also use Rust to develop a web API, but it wasn’t designed with this use case in mind. Rust’s focus on memory-safety increases complexity and development time, especially for a fairly simple web API. Programming languages that support concurrency better help you to build parallel systems. Concurrency reduces the idle time of a computer system, therefore, languages that support concurrency are popular. Go supports concurrency better than Rust. Your chances of successfully building parallel systems improve if you use Go. Go is easy to understand and use. Rust has more powerful features, which also increases its complexity. Go is an easy-to-use language, therefore, you can develop apps faster with it. On the other hand, the complexities of Rust can slow down the development. Rust has more powerful features than Go. You can achieve more with Rust while coding less.

The Third suggestion for backend development in 2021 is Kotlin

The Fourth suggestion for backend development in 2021 is NodeJS

Justin O'Brien from Competitive team on Valorent! said that the entire backend microservice architecture is built using Golang. This means that everything from spinning up and managing a game server process to purchasing items is all done using services written in golang.

## Variable-and-loop

Variables are just pointing to a memory location. `var i int` here `i` is a variable(name of the variable) which pointing a memory location. `int` is a data type which is used for telling the compiler how much memory is required to store data pointing by `i` variable.
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. `i:=10` here `i` is the name of the variable and data is `10` compiler automatically figureout data type or in other words size of the data `10`.
size of the data type int: 
```go

int8 	8 bits      
int16 	16 bits 	
int32 	32 bits 	
int64 	64 bits 	
int 	Platform dependent
```
Every Go program must be a part of some package. A package is nothing but a directory with some code files. The package declaration must be the first line of code in your Go source file.
The purpose of a package is to design and maintain a large number of programs by grouping related features together into single units so that they can be easy to maintain and understand and independent of the other package programs. This modularity allows them to share and reuse. Go programs start running in the main package. It is a special package that is used with programs that are meant to be executable. The main() function is a special function that is the entry point of an executable program.

 `import "fmt" ` fmt is a core library package that contains functionalities related to formatting and printing output or reading input from various I/O sources. It exports functions like Println(), Printf(), Scanf() etc, for other packages to reuse.

In GoLang a variable can be accessible from 3 scope packageLevelScope - accessiblee from only the specific package it defined, GlobalLevelScope - accessiblee from any packages, blockLevelScope - accessiblee from only the specific block-{ } it defined.

```go
var i int // // SYNTAX IS CORRECT
//Here i is the name of the variable. Type of the variable is int. 

package main
import "fmt" 

var packageLevelScope int =12 // this variable accessible from the main package
var GlobalLevelScope int =13 // this variable accessible from any packages
func main() {
	var blockLevelScope int=15 // this variable only accessible from main function
	k := 3 
	m, n, p, q := 1, 34.67, true, "i am string"
	// Outside a function, every statement begins with a keyword (var, func, and so on)
	// and so the := shorthand construct is not available.
	a:=10 // binary 1010
	b:=3  // binary 0011
	a&b   // 0010 -> 2 [1-1=1, 1-0=0, 0-0=0, 0-1=0]
	a|b   // 1011 -> 11 [1-1=1, 1-0=1, 0-0=0, 0-1=1]
	a^b   // 1001 -> 9 [1-1=0, 1-0=1, 0-0=0, 0-1=1]
	a&^b  // 0100 -> 8 [1-1=0, 1-0=0, 0-0=1, 0-1=0]


// FOR loop TYPE 1

    // Here for loop has 3 statements, First is initializer where k is initialized with 0 by shorthand construct :=,
    // Second is condition(comaprison) where k<10 is tell the compiler that executes the body of the loop until k<10 is became false. When k becames 10 loop will terminate.
    // Third is increment/decrement which is responsible for make the condition part false at a time. 

	for k:=0;k<10;k++{
		if k%2==0{
            continue // Continue statement at first break the execution of the block, then again start the block for execution.
            //skips the rest of the loop body and starts the next iteration.
		}
		fmt.Print(k) 
    }// 13579
    


	for row:=1; row<=3;row++{
		for col:=1;col<=3;col++{
			fmt.Print(row,":",col," -> ")
			if row*col>=3{
                break  // break statement break the block for execution
                //stops further execution of a loop construct.
			}
		}
	} // 1:1 -> 1:2 -> 1:3 -> 2:1 -> 2:2 -> 3:1 -> 

// Label is just refer a point, we want to point by the execution pointer, from where we want to start the execution.

// goto statement is used to alter the normal execution flow and transfer control to a labeld statement in the same program.
// when goto statement is encountered, compiler transfer the control to a label: and start execution from there. But I personally dont prefer to use goto in programming.

MyLabel: 
	for row:=1; row<=3;row++{
		for col:=1;col<=3;col++{
			fmt.Print(row,":",col," -> ")
			if row*col>=3{
				break MyLabel
			}
		}
	} // 1:1 -> 1:2 -> 1:3 ->


// FOR loop TYPE 2

    // index variable contains index and value variable contains each charachter at a time of a string.
//  here characters are transformed to is ascii value so we need type casting.

    s:="Rupam Ganguly"
	for index,value:=range s{
		fmt.Print(index," :- ", value," -> ")
	}// 0 :- 82 -> 1 :- 117 -> 2 :- 112 -> 3 :- 97 -> 4 :- 109 -> 5 :-  32 -> 6 :- 71 -> 7 :- 97 -> 8 :- 110 -> 9 :- 103 -> 10 :- 117 -> 11 :- 108 -> 12 :- 121 ->


	s:="Rupam Ganguly"
	for index,value:=range s{
		fmt.Print(index," :- ", string(value)," -> ") // Convert or typecast from int to string
    } // 0 :- R -> 1 :- u -> 2 :- p -> 3 :- a -> 4 :- m -> 5 :-   -> 6 :- G -> 7 :- a -> 8 :- n -> 9 :- g -> 10 :- u -> 11 :- l -> 12 :- y -> 
    
// FOR loop TYPE 3
    // here initialization step is performed at package level , only condition step is present in for loop,
    // increment step performs from body of the for loop.

	i:=1
	for ;i<10; {
		fmt.Print(i)
		i++
	}  // 123456789
	fmt.Println("")
	fmt.Println(i)// 10


// FOR loop TYPE 4
// WHILE loop of Golang
	j:=1
	for j<6{
	fmt.Print(j)
	i++
	}// 12345

 }

```
## Function-and-Panic

Syntax of Function in golang:

```go
func function_name (argument1 type, argument2 type) (return_type1, return_type2){
	return argument1,argument2
}
```
A method is nothing but a function, but it belongs to a certain type. A method is defined with slightly different syntax than a normal function. It required an additional parameter known as a receiver which is a type to which the function belongs. This way, a method (function) can access the properties of the receiver it belongs to (like fields of a struct).

When a method belongs to a type, its receiver receives a copy of the object on which it was called. 

```go
func (recieveCopy type)function_name (argument1 type, argument2 type) (return_type1, return_type2){
	return argument1,argument2
}
```
A method can also belong to the pointer of a type. When a method belongs to the pointer of a type, its receiver will receive the pointer to the object instead of a copy of the object.
If a method has a pointer receiver, then you don’t necessarily need to use the pointer dereferencing syntax (*e) to get the value of the receiver. 

```go
func (recieveReference *type)function_name (argument1 type, argument2 type) (return_type1, return_type2){
	return argument1,argument2
}
```
Panic is a built-in function that stops the ordinary flow of control and begins panicking.
The argument passed to the panic function will be printed when the program terminates.
 When the function lets Fu function calls panic, execution of Fu stops, any deferred functions in Fu are executed normally, and then Fu returns to its caller. 

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.


```go
package main

import "fmt"

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	}
	return 1
}
func anotherFunc() {
	fmt.Println("I am Another Function")
}

// Function as value - Anonymous function
// its nothing but just the name of the function is not present, and then we srore the return of the function, to a variable result.
// we can call it via `result(6, 9)`

var result = func(a, b int) int {
	return a + b
}

// Aother example: 

// subtra := func(a, b int) int {
	// 	return a - b
	// }(89, 20)

// now we can use only the variable name(subtra), as arguments(89,20) are passed at the time of definition 



func main() {
	fmt.Println("Factorial of 4", factorial(4)) //Factorial of 4 24
    defer anotherFunc()                         //defer keword in GO makes a function execute at the end of the execution (or when hits return statement) of parent function from where it is called.
	
	fmt.Println("HI I am MAIN and I execute first instead of defer function")
	fmt.Println("HI I am MAIN and I execute first instead of defer function")
	fmt.Println("HI I am MAIN and I execute first instead of defer function")
	fmt.Println("Calling Function as value - Anonymous function", result(6, 9))
	subtra := func(a, b int) int {
		return a - b
	}(89, 20)
    fmt.Println("Subtra is caling: ", subtra)
    
	defer fmt.Print("First	")
	defer fmt.Print("Second	")
	defer fmt.Print("Third	")


	fmt.Println("---Start")
	defer fmt.Println("---This is defered")

	fmt.Println("---End")

    panicker() // execution does not stops from here as recover() is present
    
    fmt.Println("+++DONE the Main")
    
	panic("BAD THING HAPPENED") // execution stops from here
	fmt.Println("I am ignored by the panic as recover() is not present for the last panic")
}



func panicker() {
	fmt.Println("About to Panic")
	defer func() {

		// synatax of another kind of if block used here:
		//  if intialization; condition{ body }
		if err := recover(); err != nil {
			fmt.Println("ERROR: ", err) // ERROR:  Yoo BAD THING HAPPEN AGAIN
		}
	}() // here `()` indicates we call the anonymous function after define it.
	panic("Yoo BAD THING HAPPEN AGAIN") // execution stops from here
	fmt.Println("DONE PANICKING") // this line never get executed
}
```
#### OUTPUT
```shell
Factorial of 4 24
HI I am MAIN and I execute first instead of defer function
HI I am MAIN and I execute first instead of defer function
HI I am MAIN and I execute first instead of defer function
Calling Function as value - Anonymous function 15
Subtra is caling:  69
---Start
---End
About to Panic
ERROR:  Yoo BAD THING HAPPEN AGAIN
+++DONE the Main
---This is defered
Third	Second	First	I am Another Function
panic: BAD THING HAPPENED

goroutine 1 [running]:
main.main()
	/tmp/sandbox739186246/prog.go:48 +0x607

```
## Struct-Class-of-GoLang

A structure has different fields of the same or different data types.  Like a class, we can create an object from this Struct 
```go
type StructName struct {
    field1 fieldType1
    field2 fieldType2
}
```

In the above syntax, StructName is a struct type while field1 and field2 are fields of data type fieldType1 and fieldType2 respectively. 

Struct gives one more ability to add meta-data to its fields. Usually, it is used to provide transformation information on how a struct field is encoded to or decoded from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you want to, either intended for another package or for your own use.
```go
	type Employee struct {
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	salary    int    `json: "salary"`
	fullTime  int    `json: "fullTime"`
	```

```go
package main

import (
	"fmt"
)

// Anonymous fields
// You can define a struct type without declaring any field names. You have to just define the field data types and Go will use the data type declarations (keywords) as the field names.

type DataStruct struct {
	string
	bool
}
type Employee struct {
	firstname string
	lastname  string
	salary    int
	fulltime  bool
}
type myStruct struct{
	foo int
}

func main() {
	// Now that we have a struct type Employee, let’s create a struct ross from it. 
	var ross Employee
	fmt.Println(ross) // {"" 0 false}
	// The zero value of a struct is a struct with all fields set to their own zero values.
	// Hence string will have zero value of ""(can’t be printed), int will have zero value of 0 and bool will have zero value of false.

	// Initialized fields of a struct
	ross.firstname = "Rupam"
	ross.lastname = "Ganguly"
	ross.salary = 45000
	ross.fulltime = true
	fmt.Println(ross) // {Rupam Ganguly 45000 true}

	// Another way to initialized fields of a struct
	bos := Employee{
		firstname: "Rintu",
		lastname:  "Ganguly",
		salary:    58000,
		fulltime:  true,
		// The comma (,) is absolutely necessary after the value assignment of the last field while creating a struct using the above syntax.
		// This way, Go won’t add a semicolon just after the last field while compiling the code.
	}
	fmt.Println(bos) // {Rintu Ganguly 58000 true}

	// Another way to initialized fields of a struct
	mos := Employee{"my mosh", "ganguly", 90000, true}
	fmt.Println(mos) // {my mosh ganguly 90000 true}

	// Anonymous struct
	// In case of an anonymous struct, we do not define any derived struct type and we create a struct by defining the inline struct type and initial values of the struct fields in the same syntax.
	monica := struct {
		age     int
		salary  int
		teacher bool
	}{
		age: 12, salary: 12345, teacher: true,
	}
	// In the above program, we are creating a struct monica without defining a derived struct type. This is useful when you don’t want to re-use a struct type.
	fmt.Println(monica) // {12 12345 true}


	// DataStruct
	samp1 := DataStruct{"Monday time", true}
	fmt.Println(samp1) // {Monday time true}
	samp1.bool = false
	fmt.Println(samp1) // {Monday time false}
	samp2 := DataStruct{"Sunday", false}
	fmt.Println(samp2) // {Sunday false}
	samp3 := samp1
	fmt.Println(samp3) // {Monday time false}

	// Nested Struct
	// A struct field can be of any data type. Hence, it is perfectly legal to have a struct field that holds another struct.
	// Hence, a struct field can have a data type that is a struct type. When a struct field has a struct value, that struct value is called a nested struct since it is nested inside a parent struct.

	type UpdatedEmployee struct {
		firstname string
		lastname  string
		salary    int
		fulltime  bool
		data      DataStruct
	}
	nestedstructMy := UpdatedEmployee{
		firstname: "Rintu",
		lastname:  "Ganguly",
		salary:    58000,
		fulltime:  true,
		data:      DataStruct{"new data inserted", true},
	}
	fmt.Println(nestedstructMy)             // {Rintu Ganguly 58000 true {new data inserted true}}
	fmt.Println(nestedstructMy.salary)      // 58000
	fmt.Println(nestedstructMy.data.string) // new data inserted

	var ms *myStruct
	ms=new(myStruct) // alternate syntax: ms=&myStruct{foo:42}
	(*ms).foo=42 // alternate syntax: ms.foo=42 // derefercencing
	fmt.Println((*ms).foo) // alternate syntax: fmt.Println((*ms).foo)

	// To export the field names of a struct, we’ve to follow the same uppercase letter approach.
	// Two structs are comparable if they belong to the same type and have the same field values.

}
```
#### OUTPUT
```shell
{  0 false}
{Rupam Ganguly 45000 true}
{Rintu Ganguly 58000 true}
{my mosh ganguly 90000 true}
{12 12345 true}
&{Rintu Ganguly 58000 true}
{Rintu Ganguly 58000 true}
first name :  Rintu
Rintu
{Monday time true}
{Monday time false}
{Sunday false}
{Monday time false}
{Rintu Ganguly 58000 true {new data inserted true}}
58000
new data inserted
42
```

## Pointer

```go
package main

import (
	"fmt"
)

func passedasReferencee(a *int) { // d is passed
	fmt.Println("*a is : ", *a)
	fmt.Println("a is : ", a)
	fmt.Println("address of a variable is: ", &a)
	fmt.Println("Incrementing:")
	*a++ // incrementing value - *a was 7 , now *a will be 8
	fmt.Println("*a is : ", *a)
	fmt.Println("a is : ", a)

}

func passedasvariable(a int) { // a is copy of c
	fmt.Println("a is : ", a)
	fmt.Println("address of a variable is: ", &a)
	fmt.Println("Incrementing:")
	a++ //a is incremented from 7 to 8
	fmt.Println("a is : ", a)
}


func main() {
	c := 7
	fmt.Println("in main function address of c is : ", &c)
	d := &c // d is memory address of c
	fmt.Println("in main function d is : ", d)
	fmt.Println("in main function *d is : ", *d)
	fmt.Println("passedasReferencee called:--->")
	passedasReferencee(d)
	fmt.Println("after executing function :- ")
	fmt.Println("c is : ", c)
	fmt.Println("d is : ", d)
	fmt.Println("*d is ", *d)
	fmt.Println("passedasVariable called:--->")
	passedasvariable(c) // here we passed c
	fmt.Println("after executing function :- ")
	fmt.Println("c is : ", c) // no change
	fmt.Println("d is : ", d)
}

```
#### 		OUTPUT:
```shell
	in main function address of c is :  0xc000012090
	in main function d is :  0xc000012090
	in main function *d is :  7
	passedasReferencee called:--->       
	*a is :  7
	a is :  0xc000012090
	address of a variable is:  0xc000006030
	Incrementing:
	*a is :  8
	a is :  0xc000012090
	after executing function :-
	c is :  8
	d is :  0xc000012090
	*d is  8
	passedasVariable called:--->
	a is :  8
	address of a variable is:  0xc000012098
	Incrementing:
	a is :  9
	after executing function :-
	c is :  8
	d is :  0xc000012090
```
## Why-use-Pointer-in-code
```go
package main

import "fmt"

var packageLevelScopeVariable int = 12

type Stack []string // the name of the array is Stack, type is string
func noReciver() {
	packageLevelScopeVariable++
	fmt.Println("noReciever->", packageLevelScopeVariable)
}
func (Stack) withReciver() {
	packageLevelScopeVariable++
	fmt.Println("withReciver->", packageLevelScopeVariable)
	//Stack = append(Stack, "one") // : type stack is not an expression
	//fmt.Println(Stack) // : type stack is not an expression

}
func (s Stack) withReciverObject_1() {
	packageLevelScopeVariable++
	fmt.Println("withReciverObject_1->", packageLevelScopeVariable)
	s = append(s, "two")
	fmt.Println(s)
	s = append(s, "three")
	fmt.Println(s)
}
func (s Stack) withReciverObject_2() {
	packageLevelScopeVariable++
	fmt.Println("withReciverObject_2->", packageLevelScopeVariable)
	s = append(s, "four")
	fmt.Println(s)
	s = append(s, "five")
	fmt.Println(s)
}
func (s *Stack) withReciverPOINTERObject_1() {
	packageLevelScopeVariable++
	fmt.Println("withReciverPOINTERObject_1->", packageLevelScopeVariable)
	*s = append(*s, "six")
	fmt.Println(s)
	*s = append(*s, "seven")
	fmt.Println(*s)
}
func (s *Stack) withReciverPOINTERObject_2() {
	packageLevelScopeVariable++
	fmt.Println("withReciverPOINTERObject_2->", packageLevelScopeVariable)
	*s = append(*s, "eight")
	fmt.Println(s)
	*s = append(*s, "nine")
	fmt.Println(*s)
}
func main() {
	noReciver()
// noReciever-> 13
	var st Stack
	st.withReciver()
// withReciver-> 14
	st.withReciverObject_1()
// withReciverObject_1-> 15
// [two]
// [two three]
	st.withReciverObject_2()
// withReciverObject_2-> 16
// [four]
// [four five]
	st.withReciverPOINTERObject_1()
// withReciverPOINTERObject_1-> 17
// &[six]
// [six seven]
	st.withReciverPOINTERObject_2()
// withReciverPOINTERObject_2-> 18
// &[six seven eight]
// [six seven eight nine]
}


```
#### OUTPUT
```shell
noReciever-> 13
withReciver-> 14
withReciverObject_1-> 15
[two]
[two three]
withReciverObject_2-> 16
[four]
[four five]
withReciverPOINTERObject_1-> 17
&[six]
[six seven]
withReciverPOINTERObject_2-> 18
&[six seven eight]
[six seven eight nine]
```
## Interface

An interface is a collection of method signatures that a Type can implement (using methods). Hence interface defines (not declares) the behavior of the object (of the type Type). The primary job of an interface is to provide only method signatures consisting of the method name, input arguments and return types. 

When an interface has zero methods, it is called an empty interface. This is represented by interface{}
```go
package main

import (
	"fmt"
	"math"
)

// in order to successfully implemented an interface,
// we need to implement all the methods declared by the interface with exact signature
type Shape interface {
	// just a struct or class with only some function declaration....
	Area() float64
	Perimeter() float64
}
type AnotherShape interface {
	// just a struct or class with only some function declaration....
	NewArea() string
	NewPerimeter() float64
}
type Rect struct {
	// this is a stuct with only some undefined variables
	width, height float64
}
type Circle struct {
	// this is a stuct with only some undefined variables
	radius float64
}


// here(r Rect) says r is object of Rect class.. 
// actually here, this is the synatax of defining a function outside of a struct
// here we define Area function and Perimeter function of Shape Interfae inside Rect struct.. so that Rect class can implements Shape interface...
// I know syntax is funny.. but its very powerful.. in this one line, compiler do lots of work
func (r Rect) Area() float64 {
	fmt.Println(r.width * r.height)
	return 7
}
func (r Rect) Perimeter() float64 {
	fmt.Println(2 * (r.width + r.height))
	return 7
}
// Now Circle class also try to implements Shape interface 
func (c Circle) Perimeter() float64 {
	fmt.Println(2 * (math.Pi * c.radius))
	return 7
}
func (c Circle) Area() float64 {
	fmt.Println(math.Pi * c.radius * 2)
	return 7
}

// Here Rect class again implements AnotherShape interface.. so Rect class implements 2 interfaces..
func (r Rect) NewArea() string {
	return "Implements multiple interfaces"
}
func (r Rect) NewPerimeter() float64 {
	return 9874321
}
func main() {
	var s Shape = Rect{10, 3}
	s.Area() // 30
	// multiple interface of Rect
	var ns AnotherShape = Rect{5, 8}
	fmt.Println(ns.NewArea())      // Implements multiple interfaces
	fmt.Println(ns.NewPerimeter()) // 9.874321e+06
	s = Circle{10}
	s.Perimeter() // 62.83185307179586

	//when Area() didnt implemented by c Circle then we get error:
	//cannot use Circle literal (type Circle) as type Shape in assignment:
	//Circle does not implement Shape (missing Area method)

}
```
#### OUTPUT
```shell
30
Implements multiple interfaces
9.874321e+06
62.83185307179586

```
Two interfaces can be compared with == and != operators. Two interfaces are always equal if the underlying dynamic values are nil, which means, two nil interfaces are always equal, hence == operation returns true.

We can find out the underlying dynamic value of an interface using the syntax i.(Type) where i is a variable of type interface and Type is a type that implements the interface. Go will check if dynamic type of i is identical to the Type and return the dynamic value is possible.

```go
package main

import "fmt"

func explain(i interface{}) {
	switch i.(type) { // we pass type of i in switch, accroding to the type cases are happen
	case string:
		fmt.Println(" Interface has a String")
	case int:
		fmt.Println("Interface has a Int")
	default:
		fmt.Println("Interface has other Type")
	}
}  
func main() {
	explain("Rupam Ganguly")
	explain(45.987)
	explain(43)
}
```
#### OUTPUT
```shell
Interface has a String
Interface has other Type
Interface has a Int
```
## Mimic-of-Inheritance

A key feature supporting traditional object oriented design is inheritance. Inheritance supports sharing of code and data, between related objects.  Inheritance means inheriting the properties of the superclass into the base class and is one of the most important concepts in Object-Oriented Programming. Since Golang does not support classes, so inheritance takes place through struct embedding. We cannot directly extend structs but rather use a concept called composition where the struct is used to form other objects. So, you can say there is No Inheritance Concept in Golang.
composition relationship ["has a" relationship]
In composition, base structs can be embedded into a child struct and the methods of the base/parent struct can be directly called on the child struct 
Multiple inheritances take place when the child struct is able to access multiple properties, fields, and methods of more than one base struct. 

```go
// In GO we can use struct for inheritance, we can compose using structs to form other objects
package main

import "fmt"

type ParentClass struct {
	color      string
	background string
}

// thats how we define function of ParentClass struct -outsude of the class/struct....
func (ParentClass) basFuncOne() string {
	return "I am Base Function ONE"
}
func (ParentClass) basFuncTwo() string {
	return "I am Base Function TWO"
}

type ChildClass struct {
	ParentClass
	style    string
	fontsize int
}


// thats how we define function of ChildClass outside class/struct
func (c ChildClass) childFunc() string {
	return "I am Child Function "
}
func main() {
	// Patent object creation
	baseObj := ParentClass{
		color:      "Black",
		background: "White",
	}
	// Child object creation
	childObj := ChildClass{
		ParentClass: baseObj,
		style:       "bold",
		fontsize:    32,
	}
	fmt.Println(childObj)    // {{Black White} bold 32}
	childObj.color = "Green" // childObj can access color backgroud fields of Base STRUCT
	childObj.background = "Red"
	childObj.fontsize = 16
	fmt.Println(childObj)              // {{Green Red} bold 16}
	fmt.Println(childObj.basFuncOne()) // childObj can access basFunc Function of Base STRUCT
	fmt.Println(childObj.basFuncTwo()) // I am Base Function ONE // I am Base Function TWO
	fmt.Println(childObj.childFunc())
}
```
#### OUTPUT
```shell
{{Black White} bold 32}
{{Green Red} bold 16}
I am Base Function ONE
I am Base Function TWO
I am Child Function 
```
### ComandLine-Argument-Passing

we can run go file with passing arguments to it, like: `go run demo.go 1234` here 1234 is the agument of int type.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// If you’re wondering why we expect 2 arguments,
		// it’s because the first argument – at index 0 – is always
		// the path of the currently running executable.
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
```
#### OUTPUT
```shell
PS E:\PROJECTS\BACKEND\gobackendcrud> go run demo.go
exit status 1
PS E:\PROJECTS\BACKEND\gobackendcrud> go run demo.go 1234
It's over 1234
PS E:\PROJECTS\BACKEND\gobackendcrud> 
```
## Allocation

Go has two allocation primitives, the built-in functions new and make. new is a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it. 
That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T. 

```go
package main

import "fmt"

type myBox struct {
	index int
	data  int
}

func main() {
	obj1 := new(myBox)
	fmt.Println(obj1)  // &{0 0}
	fmt.Println(*obj1) // {0 0}
	var obj2 myBox
	fmt.Println(obj2) // {0 0}
	//fmt.Println(*obj2) // invalid indirect of obj2 (type myBox)
}
```
#### OUTPUT
```shell
&{0 0}
{0 0}
{0 0}
```
The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use. Example of make function is present in Slice section...
<style>
body{
overflow-wrap: normal;
color:red;
}
</style>
