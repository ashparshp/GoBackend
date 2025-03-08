# Go Notes

## 1. Basic Printing (`fmt` Package)

### `fmt.Print()`

- Prints output **without a newline**.
- Multiple arguments are concatenated with spaces.

```go
package main
import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Print("World!") // Output: Hello, World!
}
```

### `fmt.Println()`

- Similar to `Print()`, but adds a **newline** at the end.

```go
fmt.Println("Hello, World!") // Output: Hello, World!\n
```

### `fmt.Printf()`

- Allows **formatted printing** using format specifiers (`%s`, `%d`, `%v`, etc.).

```go
name := "Alice"
age := 25
fmt.Printf("My name is %s and I am %d years old.\n", name, age)
// Output: My name is Alice and I am 25 years old.
```

---

## 2. Formatted String (`fmt.Sprintf()`)

- Instead of printing, it **returns** a formatted string.

```go
message := fmt.Sprintf("Hello, %s!", "Alice")
fmt.Println(message) // Output: Hello, Alice!
```

---

## 3. Logging Output (`log` Package)

### `log.Print()`, `log.Println()`, `log.Printf()`

- Works similarly to `fmt`, but adds timestamps and is **better for logging**.

```go
import "log"

log.Println("This is a log message.")
log.Printf("User %s has logged in.", "Alice")
```

### `log.Fatal()` and `log.Panic()`

- `log.Fatal()`: Prints the message and **exits the program**.
- `log.Panic()`: Prints the message and **panics** (throws an error).

```go
log.Fatal("Something went wrong!") // Exits program after printing.
log.Panic("Something went wrong!") // Prints and panics.
```

---

## 4. Writing to Standard Output (`os.Stdout`)

- Can be used for **lower-level** output control.

```go
import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stdout, "Hello from os.Stdout!")
}
```

- For error messages:

```go
fmt.Fprintln(os.Stderr, "An error occurred!")
```

---

## 5. `println()` (Built-in Function)

### 🔹 What is `println()`?

- A **built-in function** (not part of `fmt`).
- Mainly for **debugging**, **not recommended** for general use.
- Arguments are separated by spaces, **no formatting support**.

#### Example:

```go
package main

func main() {
    println("Hello", "World", 42)
}
```

**Output:**

```
Hello World 42
```

### ⚠️ Why Avoid `println()`?

1. **No formatting support** (unlike `fmt.Printf`).
2. **Implementation-dependent** (not consistent across platforms).
3. **Not recommended** for general use (only for debugging).

### ✅ Preferred Alternative: `fmt.Println()`

```go
import "fmt"

func main() {
    fmt.Println("Hello", "World", 42)
}
```

---

## 🔥 Key Differences: `println()` vs. `fmt.Println()`

| Feature              | `println()` (Built-in) | `fmt.Println()`        |
| -------------------- | ---------------------- | ---------------------- |
| Package              | Built-in (no import)   | `fmt` (needs import)   |
| Formatting           | ❌ No support          | ✅ Supports formatting |
| Output Stream        | Varies                 | Uses `os.Stdout`       |
| Usage Recommendation | ❌ Not recommended     | ✅ Recommended         |

---

## 6. Go Build & Run Commands

### 🛠️ Run Go Code Without Building an Executable

```sh
go run main.go
```

- Compiles and runs the file without creating an executable.

### 🏗️ Compile and Build an Executable

```sh
go build main.go
./main        # (Linux/macOS)
main.exe      # (Windows)
```

- Generates an executable in the current directory.

### 📦 Install a Go Binary

```sh
go install
```

- Builds and installs the binary in `$GOPATH/bin`.

### 🌍 Cross-Compile for Different OS/Architectures

```sh
GOOS=windows GOARCH=amd64 go build -o myapp.exe main.go  # Windows 64-bit
GOOS=linux GOARCH=amd64 go build -o myapp main.go       # Linux 64-bit
GOOS=darwin GOARCH=arm64 go build -o myapp main.go      # macOS M1/M2
```

### 🔍 Verbose Mode for Debugging

```sh
go build -v
```

```sh
go run -v main.go
```

### 🧹 Clean Build Files

```sh
go clean
```

```sh
go clean -modcache
```

---

## 7. Go Variables & Values

### 🏷️ Declaring Variables

```go
var name string = "Rahul"
var age int = 25
var height float32 = 5.8
var isStudent bool = true
fmt.Println(name, age, height, isStudent)
```

### 🔹 Multiple Variable Declaration

```go
var (
    country string = "USA"
    city    string = "New York"
)
fmt.Println(country, city)
```

### ⚡ Constants

```go
const pi = 3.14
fmt.Println(pi)
```

### 🛠️ Default Zero Values

```go
var a int    // Default: 0
var b string // Default: ""
var c bool   // Default: false
fmt.Println(a, b, c)
```

### 🎯 Type Inference

```go
var n = 10 // Type inferred as int
fmt.Println(n)
```

### ⚡ Short Variable Declaration (`:=`)

```go
firstName := "Rahul"
fmt.Println(firstName)
```

---

## 8. Loops in Go

### 🔄 `for` Loop (Basic)

```go
for i := 0; i < 5; i++ {
    fmt.Println("Iteration", i)
}
```

### 🔁 `for` as a While Loop

```go
i := 0
for i < 5 {
    fmt.Println("Iteration", i)
    i++
}
```

### 🔄 Infinite Loop

```go
for {
    fmt.Println("This will run forever!")
}
```

### 🛑 Breaking Out of a Loop

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}
```

### 🔄 Skipping Iterations with `continue`

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue // Skips iteration when i == 2
    }
    fmt.Println(i)
}
```

### 🔄 Looping with `range`

```go
for i := range 11 {
    fmt.Println(i)
}
```

```go
numbers := []int{10, 20, 30, 40}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

```go
word := "Golang"
for _, char := range word {
    fmt.Printf("%c ", char)
}
// Output: G o l a n g
```

---

## 9. Conditions in Go

### ✅ `if` Statement

```go
if age > 18 {
    fmt.Println("Adult")
}
```

### ✅ `if-else` Statement

```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### ✅ `if-else if-else` Ladder

```go
if age < 13 {
    fmt.Println("Child")
} else if age < 18 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Adult")
}
```

### ✅ Short Variable Declaration in `if`

```go
if num := 10; num%2 == 0 {
    fmt.Println("Even")
} else {
    fmt.Println("Odd")
}
```

### ✅ `switch` Statement

```go
switch day := "Monday"; day {
case "Monday":
    fmt.Println("Start of the week!")
case "Friday":
    fmt.Println("Weekend is near!")
default:
    fmt.Println("Just another day.")
}
```

### ✅ `switch` Without Condition (Acts Like `if-else`)

```go
num := 10
switch {
case num < 0:
    fmt.Println("Negative")
case num == 0:
    fmt.Println("Zero")
default:
    fmt.Println("Positive")
}
```

### ✅ `switch` case with multiple cases

```go
switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
}
```

### ✅ `Type` switch

```go
whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("Unknown type")
		}
	}
	whoAmI(1)
	whoAmI("Hello")
	whoAmI(1.0)
```

---

## 11. Arrays in Go

### Declaring an Array

- Arrays have a fixed size and store multiple values of the same type.

```go
var numbers [5]int
fmt.Println(numbers)       // [0 0 0 0 0]
fmt.Println(len(numbers))  // 5
```

### Initializing an Array

- Explicitly specifying values.

```go
numbers := [5]int{1, 2, 3, 4, 5}
fmt.Println(numbers) // [1 2 3 4 5]
```

### Implicit Length Deduction

- The compiler determines the array length.

```go
numbers := [...]int{1, 2, 3, 4, 5}
fmt.Println(numbers) // [1 2 3 4 5]
```

### Index-based Initialization

```go
numbers := [5]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
fmt.Println(numbers) // [1 2 3 4 5]
```

### 2D Arrays

```go
matrix := [2][2]int{{1, 2}, {3, 4}}
fmt.Println(matrix)       // [[1 2] [3 4]]
fmt.Println(matrix[0])    // [1 2]
fmt.Println(matrix[1])    // [3 4]
fmt.Println(matrix[0][1]) // 2
```

## 11. Slices in Go

### Uninitialized Slice

- An uninitialized slice is `nil`.
- Its length is `0`.

```go
var nums []int
fmt.Println("nums:", nums)
fmt.Println("Is nil:", nums == nil)
fmt.Println("Length:", len(nums))
```

### Initialized Slice

- Created using `make([]T, length, capacity)`.
- Length: Number of elements initially present.
- Capacity: Maximum storage before resizing.

```go
var nums2 = make([]int, 2, 5)
fmt.Println("nums2:", nums2)
fmt.Println("Is nil:", nums2 == nil)
fmt.Println("Length:", len(nums2))
fmt.Println("Capacity:", cap(nums2))
```

### Appending Elements

- `append()` adds elements, growing capacity dynamically.

```go
nums2 = append(nums2, 1, 4, 5, 6)
fmt.Println("After appending elements:", nums2)
fmt.Println("New Length:", len(nums2))
fmt.Println("New Capacity:", cap(nums2))
```

### Slice Literal

- A slice can be initialized directly with values.

```go
nums3 := []int{}
nums3 = append(nums3, 1, 2)
fmt.Println("nums3:", nums3)
fmt.Println("Length:", len(nums3))
fmt.Println("Capacity:", cap(nums3))
```

### Copying a Slice

- Use `copy(destination, source)` to duplicate a slice.

```go
nums4 := make([]int, len(nums3))
copy(nums4, nums3)
fmt.Println("nums4 (copied from nums3):", nums4)
```

### Slice Operator

- Extracts sub-slices.

```go
var nums5 = []int{1, 2, 3, 4, 5}
fmt.Println("nums5[1:3]:", nums5[1:3]) // [2 3]
fmt.Println("nums5[:3]:", nums5[:3])   // [1 2 3]
fmt.Println("nums5[1:]:", nums5[1:])   // [2 3 4 5]
fmt.Println("nums5[:]:", nums5[:])     // [1 2 3 4 5]
```

### Slice Functions

- `slices.Equal()` compares slices.

```go
var nums6 = []int{1, 2, 3, 4, 5}
var nums7 = []int{6, 7, 8, 9, 10}
fmt.Println("Are nums6 and nums7 equal?", slices.Equal(nums6, nums7)) // false
fmt.Println("Are nums6 and nums6 equal?", slices.Equal(nums6, nums6)) // true
```

### 2D Slice

- A slice of slices.

```go
var nums8 = [][]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println("nums8:", nums8)
fmt.Println("First row:", nums8[0])
fmt.Println("Element at [0][1]:", nums8[0][1])
```

### Iterating Over a 2D Slice

```go
for i, v := range nums8 {
    fmt.Println("Row", i, ":", v)
}
```

# 12. Go Maps

## Overview

Maps in Go are built-in associative data structures, similar to dictionaries in Python or hashes in Ruby. They store key-value pairs and provide efficient lookups.

---

## Creating a Map

You can create a map using `make()` or a map literal.

```go
// Using make()
m := make(map[string]int) // Creates an empty map with string keys and int values

// Using a map literal
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n)
```

## Adding and Retrieving Values

```go
m["k1"] = 7
m["k2"] = 13
fmt.Println("map:", m) // Output: map[k1:7 k2:13]

// Retrieving values
v1 := m["k1"]
fmt.Println("v1: ", v1) // Output: v1: 7
```

## Getting the Length of a Map

```go
fmt.Println("len:", len(m)) // Output: length of map
```

## Deleting a Key from a Map

```go
delete(m, "k2")
fmt.Println("map:", m) // k2 is removed
```

## Clearing All Key/Value Pairs

```go
clear(m)
fmt.Println("map:", m) // Output: map: {}
```

## Checking if a Key Exists

```go
value, exists := m["k2"]
if exists {
    fmt.Println("Key exists, value:", value)
} else {
    fmt.Println("Key not found")
}
```

## Iterating Over a Map

```go
for key, value := range n {
    fmt.Println("key:", key, "value:", value)
}
```

## Comparing Two Maps

Since maps are reference types, you cannot compare them directly using `==`. However, you can use the `maps.Equal()` function (introduced in Go 1.21) to compare two maps.

```go
import "maps"

m1 := map[string]int{"foo": 1, "bar": 2}
m2 := map[string]int{"foo": 1, "bar": 2}
fmt.Println("m1 == m2:", maps.Equal(m1, m2)) // Output: true
```

## Notes

- A key that does not exist in the map returns the zero value of the value type.
- Maps are **not safe for concurrent use**; use `sync.Mutex` or `sync.Map` for safe concurrent access.
- The `delete()` function does nothing if the key does not exist.
- The order of iteration over a map is **not guaranteed**.

---

## 12. Functions

### Basic Function

A function in Go can take zero or more arguments and return values.

```go
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

### Omitting Types

When two or more consecutive parameters share a type, you can omit the type from all but the last.

```go
func add2(x, y int) int {
	return x + y
}
```

### Multiple Return Values

Functions can return multiple values.

```go
func swap(x, y string) (string, string, bool) {
	return y, x, true
}

func main() {
	a, b, _ := swap("hello", "world")
	fmt.Println(a, b)
}
```

### Passing Function as Argument

A function can be passed as an argument to another function.

```go
func applyIt(afunc func(int) int, val int) int {
	return afunc(val)
}

func double(n int) int {
    return n * 2
}

func main() {
	applyIt0 := applyIt(func(x int) int { return x * x }, 3)
	fmt.Println(applyIt0)

	applyIt1 := applyIt(double, 3)
	fmt.Println(applyIt1)
}
```

### Returning a Function

A function can return another function.

```go
func makeAdder() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

func main() {
	adder := makeAdder()
	fmt.Println(adder(3))
}
```

### Variadic Functions

A function can take a variable number of arguments.

```go
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Result:", result)

	nums := []int{1, 2, 3, 4, 5}
	result = sum(nums...)
	fmt.Println("Result:", result)
}
```

### Variadic Functions with Any Type

A variadic function can accept arguments of any type using `interface{}`.

```go
func sumAnyType(items ...interface{}) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	sumAnyType(1, "Hello", 3.14, true)
}
```

### Closures in Go

Closures are functions that **capture** and **remember** the variables in their surrounding scope.

#### Example: Basic Closure

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

increment := counter()
fmt.Println(increment()) // 1
fmt.Println(increment()) // 2
fmt.Println(increment()) // 3
```

#### How Closures Work

1. The `counter` function defines a local variable `count`.
2. It returns an **anonymous function** that modifies and returns `count`.
3. Each call to `increment()` retains access to `count`.

#### Example: Multiple Closures

Each call to `counter()` creates a **new** closure with its own `count`.

```go
inc1 := counter()
inc2 := counter()

fmt.Println(inc1()) // 1
fmt.Println(inc1()) // 2
fmt.Println(inc2()) // 1 (separate counter)
```

Closures allow state to be maintained across multiple function calls while avoiding global variables.

---

## 13. Pass by Value vs Pass by Reference

Go typically uses **pass by value** for function arguments, but you can achieve **pass by reference** using pointers.

### Pass by Value

By default, Go passes variables by value, creating copies of the original values:

```go
package main

import "fmt"

// By default, Go uses pass by value
func changeNum(num int) {
	num = 5
	fmt.Println("num in changeNum:", num)
}

func main() {
	num := 1

	// Pass by value
	fmt.Println("num before changeNum:", num)
	changeNum(num)
	fmt.Println("num after changeNum:", num)  // Original value remains unchanged
}
```

**Output:**

```
num before changeNum: 1
num in changeNum: 5
num after changeNum: 1
```

### Pass by Reference (Using Pointers)

To modify the original value, use pointers:

```go
package main

import "fmt"

// Pass by reference using pointers
func changeNumByReference(num *int) {
	*num = 5
	fmt.Println("num in changeNumByReference:", *num)
}

func main() {
	num := 1

	// To pass by reference, use pointers
	fmt.Println("num before changeNumByReference:", num)
	changeNumByReference(&num)
	fmt.Println("num after changeNumByReference:", num)  // Original value is modified
}
```

**Output:**

```
num before changeNumByReference: 1
num in changeNumByReference: 5
num after changeNumByReference: 5
```

---

## 14. Structs in Go

Structs are composite data types that group together variables under a single name.

### Defining a Struct

```go
package main

import (
	"fmt"
	"time"
)

// Define a struct
type Order struct {
	Id         int
	Amount     float64
	Status     string
	CreateDate time.Time // Nanosecond precision
	Product    string
}
```

### Creating Struct Instances

There are multiple ways to create struct instances:

```go
// Method 1: With field names
order := Order{
    Id:         1,
    Amount:     100.0,
    Status:     "pending",
    CreateDate: time.Now(),
    Product:    "Laptop",
}

// Method 2: Without field names (must follow the same order as defined)
order2 := Order{2, 200.0, "pending", time.Now(), "Mobile"}
```

### Accessing and Modifying Struct Fields

```go
// Accessing fields
fmt.Println(order.Id)
fmt.Println(order.Amount)
fmt.Println(order.Status)
fmt.Println(order.CreateDate)
fmt.Println(order.Product)

// Modifying fields
order.Status = "completed"
fmt.Println(order.Status)
```

### Receiver Functions (Methods)

Methods are functions with a special receiver argument.

```go
// Value receiver - doesn't modify the original struct
func (o Order) printOrder() {
	fmt.Println(o.Id)
	fmt.Println(o.Amount)
	fmt.Println(o.Status)
	fmt.Println(o.CreateDate)
	fmt.Println(o.Product)
}

// Pointer receiver - can modify the original struct
func (o *Order) updateStatus(status string) {
	o.Status = status
}
```

### Factory Functions

Factory functions create and return new instances of structs:

```go
// Factory function
func newOrder(id int, amount float64, status string, createDate time.Time, product string) Order {
	return Order{id, amount, status, createDate, product}
}

// Usage
order3 := newOrder(3, 300.0, "pending", time.Now(), "Tablet")
order4 := newOrder(4, 400.0, "pending", time.Now(), "Desktop")
```

### Anonymous Structs

Anonymous structs are defined and used without a separate type declaration:

```go
// Anonymous struct
language := struct {
    Name    string
    Version float64
}{
    Name:    "Go",
    Version: 1.15,
}
fmt.Println(language.Name)
fmt.Println(language.Version)
```

### Key Points

1. **Value vs Pointer Receivers**:

   - Use value receivers (`func (o Order)`) when you don't need to modify the struct
   - Use pointer receivers (`func (o *Order)`) when you need to modify the struct
   - Pointer receivers are more efficient for large structs

2. **Struct Initialization**:

   - Named fields can be in any order
   - Unnamed fields must follow the struct definition order
   - Unspecified fields get zero values

3. **Anonymous Structs**:

   - Useful for one-time use structures
   - Good for small, localized data groupings

4. **Factory Functions**:
   - Create consistent instances
   - Can implement validation logic
   - Improve code readability
