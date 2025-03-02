package main

import "fmt"

func main() {

	/* Loops*/

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	value := 0
	for value < 5 {
		fmt.Println(value)
		value++
	}

	// for loop with range
	for i := range 5 {
		fmt.Println(i)
	}

	/* Infinite Loop
	for {
		fmt.Println(1)
	}
	*/

	/* Conditional Statements*/

	// if else
	if 1 > 2 {
		fmt.Println("1 is greater than 2")
	} else {
		fmt.Println("1 is not greater than 2")
	}

	// if else if
	if 1 > 2 {
		fmt.Println("1 is greater than 2")
	}
	if 1 < 2 {
		fmt.Println("1 is less than 2")
	}

	// switch case
	switch 1 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

}
