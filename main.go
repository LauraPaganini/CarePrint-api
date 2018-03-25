package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	//create channels to catch CTRL+C to shutdown server
	sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/login/create", CreateAccountHandler)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		// The program will wait here until it gets the
		// expected signal and then exit.
		<-sigs
		fmt.Println("shutdown")
		ctx := context.Background()
		srv.Shutdown(ctx)
		close()
		os.Exit(3)
	}()

	fmt.Println("serving...")
	log.Fatal(srv.ListenAndServe())
}
