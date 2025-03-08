package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter()
	increment1 := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment1())
}