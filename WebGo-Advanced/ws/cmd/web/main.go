package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting web server on port 8081")

	_ = http.ListenAndServe(":8081", mux)
}
