package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rosariop/go-web/authentication"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials UserCredentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCredentials.Username == "someUsername" && userCredentials.Password == "somePassword" {
		exp := time.Now().Add(10 * time.Minute)
		jwt, err := authentication.GenerateJwt(userCredentials.Username, "admin", exp)
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
