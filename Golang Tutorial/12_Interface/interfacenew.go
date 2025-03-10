/*
package main

import "fmt"

type Payment struct{}

func (p Payment) makePayment(amount float64) {

	// Razorpay
	// razorPayPaymentGw := Razorpay{}
	// razorPayPaymentGw.Pay(amount)

	// Stripe
	stripePaymentGw := Stripe{}
	stripePaymentGw.Pay(amount)
}

type Razorpay struct {}

func (r Razorpay) Pay (amount float64) {
	fmt.Println("Paying using Razorpay", amount)
}

type Stripe struct {}

func (s Stripe) Pay (amount float64) {
	fmt.Println("Paying using Stripe", amount)
}

func main() {

	// paymentOne := Payment{}
	// paymentOne.makePayment(100)

	paymentTwo := Payment{}
	paymentTwo.makePayment(200)

}

/*
package main

import "fmt"

type Payment struct{
	// gateway Stripe
	gateway Razorpay
}

func (p Payment) makePayment(amount float64) {
	p.gateway.Pay(amount)
}

type Razorpay struct {}

func (r Razorpay) Pay (amount float64) {
	fmt.Println("Paying using Razorpay", amount)
}

type Stripe struct {}

func (s Stripe) Pay (amount float64) {
	fmt.Println("Paying using Stripe", amount)
}

func main() {

	// Stripe
	// stripePaymentGw := Stripe{}
	// newPayment := Payment{
	// 	gateway: stripePaymentGw,
	// }

	// Razorpay
	newPayment := Payment{
		gateway: Razorpay{},
	}

	newPayment.gateway.Pay(100)

}

*/

package main

import "fmt"

/* Paymenter Interface */
type Paymenter interface {
	Pay(amount float64)
	Refund(amount float64, account string)
}

/* Payment */
type Payment struct {
	gateway Paymenter
}

/* Razorpay */
type Razorpay struct{}

func (r Razorpay) Pay(amount float64) {
	fmt.Println("Paying using Razorpay", amount)
}

func (r Razorpay) Refund(amount float64, account string) {
	fmt.Println("Refunding using Razorpay", amount, account)
}

/* Stripe */
type Stripe struct{}

func (s Stripe) Pay(amount float64) {
	fmt.Println("Paying using Stripe", amount)
}

func (s Stripe) Refund(amount float64, account string) {
	fmt.Println("Refunding using Stripe", amount, account)
}

/* Testing */
type TestPayment struct{}

func (t TestPayment) Pay(amount float64) {
	fmt.Println("Paying using TestPayment", amount)
}

func (t TestPayment) Refund(amount float64, account string) {
	fmt.Println("Refunding using TestPayment", amount, account)
}

func main() {

	/* TestPayment */
	newPayment := Payment{
		gateway: TestPayment{},
	}
	newPayment.gateway.Pay(100)

	/* Razorpay */
	newPayment = Payment{
		gateway: Razorpay{},
	}
	newPayment.gateway.Pay(100)

	/* Stripe */
	newPayment = Payment{
		gateway: Stripe{},
	}
	newPayment.gateway.Pay(100)

	/* Refund */
	newPayment.gateway.Refund(100, "123456789")
}
