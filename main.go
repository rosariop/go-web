package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rosariop/go-web/authentication"
)

func main() {
	http.HandleFunc("/authenticate", authentication.AuthHandler)

	fmt.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
