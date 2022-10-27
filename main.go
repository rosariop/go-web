package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func generateJwt(username string, role string) (string, error) {
	secret := []byte("mysecret")

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username
	claims["roles"] = []string{role}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := at.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials UserCredentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCredentials.Username == "someUsername" && userCredentials.Password == "somePassword" {
		jwt, err := generateJwt(userCredentials.Username, "admin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		returnString := "Bearer " + jwt

		w.Write([]byte(returnString))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func main() {

	http.HandleFunc("/authenticate", authHandler)

	fmt.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
