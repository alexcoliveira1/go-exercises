package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"answer"
)

func main() {
	ctx := context.Background()
	// our answer service
	srv := answer.NewService()

	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := answer.Endpoints{
		GetHelloEndpoint: answer.MakeGetHelloEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("napodate is listening on port: 8090")
		http.Handle("/", answer.NewRestHTTPServer())
		http.Handle("/private", answer.NewGRPCHTTPServer(ctx, endpoints))
		errChan <- http.ListenAndServe(":8090", nil)
	}()

	log.Fatalln(<-errChan)
}
