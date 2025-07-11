package main

import "fmt"

func main() {

	// Variables
	var name string = "Rahul"
	var age int = 25
	var height float32 = 5.8
	var isStudent bool = true
	fmt.Println(name, age, height, isStudent)

	// Multiple Variables
	var (
		country string = "USA"
		city    string = "New York"
	)
	fmt.Println(country, city)

	// Constants
	const pi = 3.14
	fmt.Println(pi)

	// Zero Values
	var a int
	fmt.Println(a)

	// Type Inference
	var b = 10
	fmt.Println(b)

	// Short Variable Declaration
	firstName := "Rahul"
	fmt.Println(firstName)
}
