package main

import (
	"encoding/json"
	"fmt"
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
		panic(err)
	}

	hash, _ := hashPassword(request.Password)
	match := checkPasswordHash(request.Password, hash)
	fmt.Println(fmt.Sprintf("login %v", match))
}

// CreateAccountHandler is called from /login/create
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request loginRequest
	err := decoder.Decode(&request)

	if err != nil {
		panic(err)
	}

	hash, _ := hashPassword(request.Password)
	fmt.Println("create")

	dbStatement := "INSERT INTO dbo.Accounts (Email, PasswordHash) VALUES (" + request.Email + ", " + hash + ");"
	modifyData(dbStatement)

	w.WriteHeader(http.StatusOK)
}
