package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	answer "github.com/alexcoliveira1/go-exercises/final_test"
)

func main() {
	ctx := context.Background()
	// our answer service
	svc := answer.NewService()
	endpoints := answer.MakeAllEndpoints(svc)

	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport
	go func() {
		log.Println("answer is listening on port: 8090")
		errChan <- http.ListenAndServe(":8090", answer.NewHTTPServer(ctx, endpoints))
	}()

	log.Fatalln(<-errChan)
}
