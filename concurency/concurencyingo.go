package main

import (
	"fmt"
	"sync"
)

var waitGroupVar = sync.WaitGroup{}

func main() {
	waitGroupVar.Add(5)
	fmt.Println("I am Main Thread")
	go counterMessage("First")

	fmt.Println("finished")
	go counterMessage("Second")
	go counterMessage("Third")
	go justrMessage()
	go counterMessage("Foyrth")

	waitGroupVar.Wait()
}

func counterMessage(msz string) {
	for i := 0; i < 7; i++ {
		fmt.Println(i, " ", msz, " -> ")
	}
	waitGroupVar.Done()
}

func justrMessage() {
	for i := 0; i < 7; i++ {
		fmt.Println(i, " ", "just message", " -> ")
	}
	waitGroupVar.Done()
}
