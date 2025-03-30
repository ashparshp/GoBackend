package main

import (
	"fmt"
	"net/http"
	"templates/pkg/handlers"
)
const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	fmt.Println("Server running on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}