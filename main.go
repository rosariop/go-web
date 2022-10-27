package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Via Web")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Via Web")
}

func main() {

	http.HandleFunc("/authenticate", authHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
