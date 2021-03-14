package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"google.golang.org/protobuf/proto"
)

type emp struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Salary int64  `json:"salary"`
}

func main() {
	// --------------------json generator-----------------
	var list []emp
	id := 0
	for i := 0; i < 8134; i++ {
		id = id + i
		name := strconv.Itoa(i)
		e1 := emp{
			ID:     strconv.Itoa(id),
			Name:   name + "Rupam Ganguly",
			Age:    26,
			Salary: 28000,
		}
		list = append(list, e1)
	}
	e, err := json.MarshalIndent(list, "", "")
	if err != nil {
		fmt.Print("ERROR")
	}

	_ = ioutil.WriteFile("test.json", e, 0644)

	// --------------Protobuf generator-----------------------
	idproto := 0
	for i := 0; i < 8134; i++ {
		idproto = idproto + i
		name := strconv.Itoa(i)
		emp := &Employee{
			ID:     strconv.Itoa(idproto),
			Name:   name + "Rupam Ganguly",
			Age:    26,
			Salary: 32345,
		}
		data, _ := proto.Marshal(emp)
		f, _ := os.OpenFile("fname",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if _, err := f.Write(data); err != nil {
			log.Println(err)
		}
	}
}
