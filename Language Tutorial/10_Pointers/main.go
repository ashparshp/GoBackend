package main

import "fmt"

// by default, Go uses pass by value
func changeNum(num int) {
	num = 5
	fmt.Println("num in changeNum:", num)
}

// pass by reference
func changeNumByReference(num *int) {
	*num = 5
	fmt.Println("num in changeNumByReference:", *num)
}

func main() {
	num := 1
	
	// pass by value
	fmt.Println("num before changeNum:", num)
	changeNum(num)
	fmt.Println("num after changeNum:", num)

	// to pass by reference, use pointers
	fmt.Println("num before changeNumByReference:", num)
	changeNumByReference(&num)
	fmt.Println("num after changeNumByReference:", num)

}