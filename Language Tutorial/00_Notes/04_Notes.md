# Understanding Mutex in Go

Mutex (Mutual Exclusion) is a synchronization primitive used to protect shared resources from concurrent access in Go. It allows only one goroutine to access the protected resource at a time, preventing race conditions.

## Key Concepts

- **Race Condition**: When multiple goroutines access and modify shared data simultaneously, leading to unpredictable results
- **Critical Section**: A section of code that accesses shared resources and must not be executed by multiple goroutines simultaneously
- **Mutex**: A lock mechanism that ensures only one goroutine can execute a critical section at a time

## Basic Mutex Operations

- `Lock()`: Acquire the lock (blocks if already locked)
- `Unlock()`: Release the lock
- Always pair `Lock()` and `Unlock()` operations

## Example: Protecting a Counter

```go
package main

import (
    "fmt"
    "sync"
)

type post struct {
    views int
    mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()
    defer p.mu.Unlock()
    p.mu.Lock()
    p.views += 1
}

func main() {
    var wg sync.WaitGroup
    myPost := post{views: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg)
    }

    wg.Wait()
    fmt.Println(myPost.views) // Safely prints 100
}
```

## How It Works

1. Each `post` has its own mutex (`mu`) to protect its `views` field
2. The `inc()` method uses `Lock()` to acquire exclusive access to the `views` field
3. After incrementing the field, it uses `Unlock()` to release the lock
4. The `defer` statement ensures the lock is released even if the function panics
5. The `WaitGroup` ensures the main function waits for all goroutines to complete

## Best Practices

### 1. Use `defer` for Unlocking

```go
func (p *post) inc() {
    p.mu.Lock()
    defer p.mu.Unlock()

    // Critical section...
    p.views += 1
}
```

This ensures the mutex is always unlocked, even if the function returns early or panics.

### 2. Keep Critical Sections Small

```go
// Good: Lock only what needs protection
func (p *post) inc() {
    // Do expensive computation here...

    p.mu.Lock()
    p.views += 1  // Only lock the critical section
    p.mu.Unlock()

    // Do more work here...
}
```

### 3. Avoid Nested Locks

```go
// Avoid this pattern (potential deadlock)
func riskyFunction() {
    mu1.Lock()
    mu2.Lock()  // What if another goroutine has locked mu2 and is waiting for mu1?

    // Critical section...

    mu2.Unlock()
    mu1.Unlock()
}
```

### 4. Consider Read-Write Mutex (RWMutex)

When your access pattern involves many reads and few writes:

```go
type post struct {
    views int
    mu    sync.RWMutex
}

func (p *post) getViews() int {
    p.mu.RLock()  // Multiple readers can acquire read lock simultaneously
    defer p.mu.RUnlock()
    return p.views
}

func (p *post) inc() {
    p.mu.Lock()  // Writers get exclusive access
    defer p.mu.Unlock()
    p.views++
}
```

## Common Mutex Pitfalls

### 1. Forgetting to Unlock

```go
// Bad: Lock without corresponding Unlock
func badFunction() {
    mu.Lock()
    if someCondition {
        return  // Oops! The mutex remains locked
    }
    // ...
    mu.Unlock()
}
```

### 2. Copying a Mutex

```go
// Bad: Copying a mutex
type BadCounter struct {
    mu    sync.Mutex
    value int
}

func main() {
    c := BadCounter{}
    c2 := c  // Copying the mutex!

    // c and c2 now have different mutexes
}
```

### 3. Deadlock

```go
// Deadlock example
func deadlockRisk() {
    mu.Lock()
    someFunction()  // If someFunction also tries to acquire the same mutex
    mu.Unlock()
}
```

## Alternatives to Mutex

1. **Channels**: Often preferred for communication between goroutines

   ```go
   // Using channels instead of mutex
   func worker(jobs <-chan int, results chan<- int) {
       for j := range jobs {
           results <- process(j)
       }
   }
   ```

2. **atomic**: For simple operations on primitive types

   ```go
   import "sync/atomic"

   var counter int64

   atomic.AddInt64(&counter, 1)  // Atomic increment
   ```

3. **sync.Once**: For one-time initialization

   ```go
   var once sync.Once
   var instance *Singleton

   func GetInstance() *Singleton {
       once.Do(func() {
           instance = &Singleton{}
       })
       return instance
   }
   ```

## When to Use Mutex

- When multiple goroutines need to modify shared state
- When you need fine-grained control over concurrent access
- When channels would be more complex or less efficient

## When to Avoid Mutex

- When communication is the primary goal (use channels)
- When simple atomic operations suffice (use atomic package)
- When you can eliminate shared state through design
