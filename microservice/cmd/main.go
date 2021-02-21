package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	mymicroservice "projects/mymicroservice/microservice"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		hA = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()
	ctx := context.Background()
	sr := mymicroservice.NewService()
	erChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		erChan <- fmt.Errorf("%s", <-c)
	}()
	endpoints := mymicroservice.Endpoints{
		GetDataEndpoint:      mymicroservice.MakeGetDataEndpoint(sr),
		StatusEndpoint:       mymicroservice.MakeStatusEndpoint(sr),
		ValidateDataEndpoint: mymicroservice.MakeValidateDataEndpoint(sr),
	}
	go func() {
		log.Println("mymicroservice is listening on port ", *hA)
		handler := mymicroservice.NewHTTPServer(ctx, endpoints)
		erChan <- http.ListenAndServe(*hA, handler)
	}()
	log.Fatalln(<-erChan)
}
