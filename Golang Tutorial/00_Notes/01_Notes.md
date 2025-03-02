# Go Printing Methods, Build Commands & Variables

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
