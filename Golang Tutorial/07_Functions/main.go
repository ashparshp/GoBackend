package main

import "fmt"

// A function can take zero or more arguments.
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add2(x, y int) int {
	return x + y
}

// Multiple results	- A function can return any number of results.
func swap(x, y string) (string, string, bool) {
	return y, x, true
}

// Passing function as argument
func applyIt(afunc func(int) int, val int) int {
	return afunc(val)
}

func double(n int) int {
    return n * 2
}

// A function can return a function.
func makeAdder() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(add2(42, 13))

	a, b, _ := swap("hello", "world")
	fmt.Println(a, b)

	applyIt0 := applyIt(func(x int) int { return x * x }, 3)
	fmt.Println(applyIt0)

	applyIt1 := applyIt(double, 3)
	fmt.Println(applyIt1)

	adder := makeAdder()
	fmt.Println(adder(3))
}
