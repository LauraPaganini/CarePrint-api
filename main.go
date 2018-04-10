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
	//create channel to catch CTRL+C to shutdown server
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.Methods("OPTIONS").HandlerFunc(corsHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/login/create", CreateAccountHandler)
	r.HandleFunc("/upload", ImageUploadHandler)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	go func() {
		// The program will wait here until it gets the
		// expected signal and then exit.
		<-sigs
		ctx := context.Background()
		srv.Shutdown(ctx)
		closeDatabaseConnection()
		os.Exit(3)
	}()

	fmt.Println("serving...")
	openDatabaseConnection()
	log.Fatal(srv.ListenAndServe())
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
}
