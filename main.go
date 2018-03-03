package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetFoo).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetFoo (w http.ResponseWriter, r *http.Request) {}