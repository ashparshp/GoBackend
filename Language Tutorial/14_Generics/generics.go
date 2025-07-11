package main

import "fmt"

/*
func printSlice(items []int) {
	for _, v := range items {
		fmt.Println(v)
	}
}
*/

/* Print slice of any type */
func printSlice[T any](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

/* Print slice of any type 
func printSlice[T interface{}](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}
*/

/* Print slice of fixed types */
func printSliceFixed[T interface{int | string | float64}](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

/* LIFO stack: int 
type Stack struct {
	element []int
}
*/


/* LIFO stack: any */
type Stack[T any] struct {
	element []T
}


func main() {
	printSlice([]int{1, 2, 3})
	printSlice([]string{"a", "b", "c"})
	printSlice([]float64{1.1, 2.2, 3.3})

	printSliceFixed([]int{1, 2, 3})
	printSliceFixed([]string{"a", "b", "c"})
	printSliceFixed([]float64{1.1, 2.2, 3.3})
	/*
	printSliceFixed([]bool{true, false}) // Error: cannot use []bool literal (type []bool) as type []T in argument to printSliceFixed
	*/


	/* LIFO stack: int
	stack := Stack{
		element: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(stack)
	*/

	/* LIFO stack: any */
	stack := Stack[int]{
		element: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(stack)

	stack2 := Stack[string]{
		element: []string{"a", "b", "c", "d", "e"},
	}
	fmt.Println(stack2)
}
