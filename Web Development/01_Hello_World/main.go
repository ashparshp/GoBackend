package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// HelloWorld is a simple HTTP handler function that writes a response to the client
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes written: %d\n", n)
}

// HomePage is a simple HTTP handler function that writes a response to the client
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page...")
}

// AboutPage is a simple HTTP handler function that writes a response to the client
func AboutPage(w http.ResponseWriter, r *http.Request) {
	addtwo := addValues(2, 2)
	fmt.Fprintf(w, "This is the About Page... and 2 + 2 is: %d", addtwo)
}

// addValues is a simple function that adds two integers and returns the sum
func addValues(x, y int) int {
	sum := x + y
	return sum
}

// DividePage is a simple HTTP handler function that writes a response to the client
func DividePage(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
}

// divideValues is a simple function that divides two float32 values and returns the result
func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("can not divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the entry point for the application
func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/home", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/divide", DividePage)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}