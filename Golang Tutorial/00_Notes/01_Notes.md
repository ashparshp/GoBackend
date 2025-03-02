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

## 10. Arrays in Go

### ✅ Declaring an Array

```go
var arr [5]int // An array of 5 integers
fmt.Println(arr) // Output: [0 0 0 0 0]
```

### ✅ Initializing an Array

```go
numbers := [3]int{1, 2, 3} // Explicit size
names := [...]string{"Alice", "Bob", "Charlie"} // Implicit size: initialize the array without specifying the length and compiler will calculate the length
fmt.Println(numbers) // Output: [1 2 3]
fmt.Println(names)   // Output: [Alice Bob Charlie]
```

### ✅ Accessing and Modifying Elements

```go
arr[0] = 10
fmt.Println(arr[0]) // Output: 10
```

### ✅ Iterating Over an Array

```go
for i, val := range numbers {
    fmt.Printf("Index %d: %d\n", i, val)
}
```

### ✅ Multi-Dimensional Arrays

```go
var matrix [2][2]int = [2][2]int{{1, 2}, {3, 4}}
fmt.Println(matrix) // Output: [[1 2] [3 4]]
```
