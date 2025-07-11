package main

import "fmt"

func main() {

	/* Arrays have a fixed size and store multiple values of the same type. */

	// Array declaration
	fmt.Println("----- Array Declaration -----")
	var numbers [5]int
	fmt.Println("Array:", numbers)
	fmt.Println("Length:", len(numbers))

	// Assign values to the array
	numbers[0] = 1
	fmt.Println("After assigning values:", numbers)

	// Declare and initialize an array
	fmt.Println("\n----- Declare and Initialize Array -----")
	var numbers2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers2:", numbers2)

	// Declare and initialize without specifying the length (compiler calculates length)
	fmt.Println("\n----- Compiler-Deduced Length -----")
	var numbers3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("numbers3:", numbers3)

	// Declare and initialize with explicit index assignment
	fmt.Println("\n----- Explicit Index Initialization -----")
	var numbers4 = [5]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	fmt.Println("numbers4:", numbers4)

	// Declare and initialize with explicit index but without specifying length
	fmt.Println("\n----- Index Initialization Without Length -----")
	var numbers5 = [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	fmt.Println("numbers5:", numbers5)

	// Declare and initialize a 2D array
	fmt.Println("\n----- 2D Array Declaration & Initialization -----")
	var numbers6 = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("numbers6:", numbers6)
	fmt.Println("First row:", numbers6[0])
	fmt.Println("Second row:", numbers6[1])
	fmt.Println("Element at [0][1]:", numbers6[0][1])

}