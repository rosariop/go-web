package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rosariop/go-web/authentication"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no stage given")
		panic("no stage given")
	}

	if os.Args[3] == "test" {
		godotenv.Load()
	}

	http.HandleFunc("/authenticate", authentication.AuthHandler)

	fmt.Println("Starting server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
