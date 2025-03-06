package main

import "fmt"

// Variadic Functions
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Variadic Functions any type
func sumAnyType(items ...interface{}) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Result:", result)

	sumAnyType(1, "Hello", 3.14, true)

	nums := []int{1, 2, 3, 4, 5}
	result = sum(nums...)
	fmt.Println("Result:", result)
}