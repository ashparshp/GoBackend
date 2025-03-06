# Go Functions

## 1. Function Basics

Functions in Go are defined using the `func` keyword. They can take zero or more arguments and return zero or more values.

### Example:

```go
func add(x int, y int) int {
    return x + y
}
```

## 2. Omitting Repeated Types

If multiple consecutive parameters share the same type, you can omit the type for all but the last.

```go
func add2(x, y int) int {
    return x + y
}
```

## 3. Multiple Return Values

Go functions can return multiple values.

```go
func swap(x, y string) (string, string, bool) {
    return y, x, true
}
```

Usage:

```go
a, b, _ := swap("hello", "world")
fmt.Println(a, b) // Output: world hello
```

## 4. Passing Functions as Arguments

Functions in Go are first-class citizens, meaning they can be passed as arguments.

```go
func applyIt(afunc func(int) int, val int) int {
    return afunc(val)
}

func double(n int) int {
    return n * 2
}
```

Usage:

```go
result := applyIt(double, 3)
fmt.Println(result) // Output: 6
```

## 5. Returning Functions

Functions can also return other functions.

```go
func makeAdder() func(int) int {
    return func(x int) int {
        return x + 2
    }
}
```

Usage:

```go
adder := makeAdder()
fmt.Println(adder(3)) // Output: 5
```

## 6. Anonymous Functions

You can define and invoke anonymous functions inline.

```go
applyIt0 := applyIt(func(x int) int { return x * x }, 3)
fmt.Println(applyIt0) // Output: 9
```

## Summary

- Functions can take parameters and return values.
- You can omit repeated types in parameter lists.
- Functions can return multiple values.
- Functions can be passed as arguments.
- Functions can return other functions.
- Anonymous functions can be used inline.

These concepts make Go functions powerful and flexible for various programming paradigms.