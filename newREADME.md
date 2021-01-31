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
```go
func (rec recieverStruct)function_name (argument1 type, argument2 type) (return_type1, return_type2){
	return argument1,argument2
}
```
```go
func (rec *recieverStruct)function_name (argument1 type, argument2 type) (return_type1, return_type2){
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
