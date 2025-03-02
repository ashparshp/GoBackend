package main

import "fmt"

func main() {

	// Array declaration
	var numbers [5]int

	// Get the length of the array
	fmt.Println(len(numbers))

	// Assign values to the array
	numbers[0] = 1

	// Print the array
	fmt.Println(numbers)

	// Declare and initialize the array
	var numbers2 = [5]int{1, 2, 3, 4, 5}

	// Print the array
	fmt.Println(numbers2)

	// Declare and initialize the array without specifying the length and compiler will calculate the length
	var numbers3 = [...]int{1, 2, 3, 4, 5}

	// Print the array
	fmt.Println(numbers3)

	// Declare and initialize the array with index
	var numbers4 = [5]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}

	// Print the array
	fmt.Println(numbers4)

	// Declare and initialize the array with index without specifying the length
	var numbers5 = [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}

	// Print the array
	fmt.Println(numbers5)


}