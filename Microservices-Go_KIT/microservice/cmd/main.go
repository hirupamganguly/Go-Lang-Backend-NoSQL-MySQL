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
		GetDataEndpoint: mymicroservice.MakeGetDataEndpoint(sr),
		StatusEndpoint:  mymicroservice.MakeStatusEndpoint(sr),
	}
	go func() {
		log.Println("mymicroservice is listening on port ", *hA)
		handler := mymicroservice.NewHTTPServer(ctx, endpoints)
		erChan <- http.ListenAndServe(*hA, handler)
	}()
	log.Fatalln(<-erChan)
}

//
// we might want a server to gracefully shutdown when it receives a SIGTERM, or a command-line tool to stop processing input if it receives a SIGINT.

// Go signal notification works by sending os.Signal values on a channel. We’ll create a channel to receive these notifications (we’ll also make one to notify us when the program can exit).

// signal.Notify registers the given channel to receive notifications of the specified signals.

// When we run this program it will block waiting for a signal. By typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print interrupt and then exit.
