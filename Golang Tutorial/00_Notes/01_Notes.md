# Go Printing Methods

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

| Feature      | `println()` (Built-in) | `fmt.Println()` |
|-------------|--------------------|----------------|
| Package     | Built-in (no import) | `fmt` (needs import) |
| Formatting  | ❌ No support | ✅ Supports formatting |
| Output Stream | Varies | Uses `os.Stdout` |
| Usage Recommendation | ❌ Not recommended | ✅ Recommended |

---

## ✅ Conclusion
- **Use `fmt.Println()` for general printing.**
- **Use `fmt.Printf()` for formatted output.**
- **Use `log` package for logging.**
- **Use `os.Stdout` and `os.Stderr` for direct output control.**
- **Avoid `println()` unless debugging.**

