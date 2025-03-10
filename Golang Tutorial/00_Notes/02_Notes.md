# Go Payment Gateway Implementation

## Evolution of the Design

### Version 1: Hardcoded Implementation

The initial implementation directly called payment gateways from within the `Payment` struct:

```go
type Payment struct{}

func (p Payment) makePayment(amount float64) {
    // Hardcoded implementation
    stripePaymentGw := Stripe{}
    stripePaymentGw.Pay(amount)

    // To switch payment providers, code changes are required:
    // razorPayPaymentGw := Razorpay{}
    // razorPayPaymentGw.Pay(amount)
}
```

**Problems:**

- Tight coupling between `Payment` and specific gateway implementations
- Changing payment providers requires code modification
- Violates Open/Closed Principle

### Version 2: Dependency Injection

Second implementation used dependency injection to improve flexibility:

```go
type Payment struct{
    gateway Razorpay  // or gateway Stripe
}

func (p Payment) makePayment(amount float64) {
    p.gateway.Pay(amount)
}
```

**Improvements:**

- Reduced coupling through dependency injection
- Payment struct delegates to injected gateway

**Remaining Issues:**

- Still tied to concrete implementations
- Must change struct type definition to use different gateways

### Version 3: Interface-based Solution

Final implementation uses the Interface Segregation Principle:

```go
type Paymenter interface {
    Pay(amount float64)
    Refund(amount float64, account string)
}

type Payment struct {
    gateway Paymenter
}
```

**Benefits:**

- Complete decoupling using interface
- Any struct implementing the `Paymenter` interface can be used
- Easy to add new payment providers without changing existing code
- Simplified testing with test doubles
- Follows Open/Closed Principle

## Payment Gateway Implementations

### Razorpay

```go
type Razorpay struct{}

func (r Razorpay) Pay(amount float64) {
    fmt.Println("Paying using Razorpay", amount)
}

func (r Razorpay) Refund(amount float64, account string) {
    fmt.Println("Refunding using Razorpay", amount, account)
}
```

### Stripe

```go
type Stripe struct{}

func (s Stripe) Pay(amount float64) {
    fmt.Println("Paying using Stripe", amount)
}

func (s Stripe) Refund(amount float64, account string) {
    fmt.Println("Refunding using Stripe", amount, account)
}
```

### TestPayment (for testing)

```go
type TestPayment struct{}

func (t TestPayment) Pay(amount float64) {
    fmt.Println("Paying using TestPayment", amount)
}

func (t TestPayment) Refund(amount float64, account string) {
    fmt.Println("Refunding using TestPayment", amount, account)
}
```

## Usage Example

```go
func main() {
    // Using TestPayment gateway
    newPayment := Payment{
        gateway: TestPayment{},
    }
    newPayment.gateway.Pay(100)

    // Switch to Razorpay gateway
    newPayment = Payment{
        gateway: Razorpay{},
    }
    newPayment.gateway.Pay(100)

    // Switch to Stripe gateway
    newPayment = Payment{
        gateway: Stripe{},
    }
    newPayment.gateway.Pay(100)

    // Refund example
    newPayment.gateway.Refund(100, "123456789")
}
```

## Design Patterns Used

1. **Strategy Pattern**: Different payment gateway implementations can be selected at runtime
2. **Dependency Injection**: Payment gateways are injected into the Payment struct
3. **Interface Segregation Principle**: Clean interface defines only required methods
4. **Open/Closed Principle**: System is open for extension but closed for modification
