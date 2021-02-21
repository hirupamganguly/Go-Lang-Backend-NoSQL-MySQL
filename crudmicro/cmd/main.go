package main

import (
	"context"
	"time"

	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"projects/mymicroservice/crudmicro"
	"syscall"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	var hA = flag.String("http", ":8080", "http listen address")
	flag.Parse()
	ctx := context.Background()

	var ctxmongo, _ = context.WithTimeout(context.Background(), 1560*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rupamganguly:MN1ntlrWNap8l4lZ@cluster0.cpwla.mongodb.net/usermanagement?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Connect(ctxmongo)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctxmongo)

	repository, err := crudmicro.NewRepo(client, logger)

	if err != nil {
		fmt.Println(err)
	}

	sr := crudmicro.NewService(repository, logger)
	erChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		erChan <- fmt.Errorf("%s", <-c)
	}()
	endpoints := crudmicro.Endpoints{
		CreateUserEndpoint:  crudmicro.MakeCreateUserEndpoint(sr),
		GetUserByIDEndpoint: crudmicro.MakeGetUserByIDEndpoint(sr),
		GetAllUsersEndpoint: crudmicro.MakeGetAllUsersEndpoint(sr),
		DeleteUserEndpoint:  crudmicro.MakeDeleteUserEndpoint(sr),
		UpdateUserEndpoint:  crudmicro.MakeUpdateUserEndpoint(sr),
	}
	go func() {
		fmt.Println("crudmicro is listening on port: ", *hA)
		handler := crudmicro.NewHTTPServer(ctx, endpoints)
		erChan <- http.ListenAndServe(*hA, handler)
	}()
	fmt.Println(<-erChan)
}
