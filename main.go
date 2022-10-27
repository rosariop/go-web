package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Via Web")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials UserCredentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("username: %s, password: %s \n", userCredentials.Username, userCredentials.Password)

}

func main() {

	http.HandleFunc("/authenticate", authHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
