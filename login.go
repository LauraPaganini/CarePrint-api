package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

	fmt.Println("Password:", request.Password)
	fmt.Println("Hash:    ", hash)

	match := checkPasswordHash(request.Password, hash)
	fmt.Println("Match:   ", match)
}
