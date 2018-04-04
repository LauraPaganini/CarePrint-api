package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler is called from /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request loginRequest
	err := decoder.Decode(&request)

	if err != nil {
		log.Fatal(err)
	}

	rows := retrieveData("SELECT PasswordHash FROM dbo.Accounts WHERE Email='" + request.Email +"';")

	defer rows.Close()

	var passwordHash string
	for rows.Next() {
		if err := rows.Scan(&passwordHash); err != nil {
			log.Fatal(err)
		}
	}

	match := checkPasswordHash(request.Password, passwordHash)

	fmt.Println("login")

	if match {
		w.Header().Set("Access-Control-Allow-Origin", "*")
    		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CreateAccountHandler is called from /login/create
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request loginRequest
	err := decoder.Decode(&request)

	if err != nil {
		log.Fatal(err)
	}

	hash, _ := hashPassword(request.Password)
	fmt.Println("create " + hash)

	dbStatement := "INSERT INTO dbo.Accounts (Email, PasswordHash) VALUES ('" + request.Email + "', '" + hash + "');"
	err = modifyData(dbStatement)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
