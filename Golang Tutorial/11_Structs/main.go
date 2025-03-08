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

// receiver function
func (o Order) printOrder() {
	fmt.Println(o.Id)
	fmt.Println(o.Amount)
	fmt.Println(o.Status)
	fmt.Println(o.CreateDate)
	fmt.Println(o.Product)
}

// receiver function
func (o *Order) updateStatus(status string) {
	o.Status = status
}

// factory function
func newOrder(id int, amount float64, status string, createDate time.Time, product string) Order {
	return Order{id, amount, status, createDate, product}
}


func main() {
	// Create an instance of the struct
	// order := Order{
	// 	Id:         1,
	// 	Amount:     100.0,
	// 	Status:     "pending",
	// 	CreateDate: time.Now(),
	// 	Product:    "Laptop",
	// }

	// // Access the fields
	// fmt.Println(order.Id)
	// fmt.Println(order.Amount)
	// fmt.Println(order.Status)
	// fmt.Println(order.CreateDate)
	// fmt.Println(order.Product)

	// // Update the fields
	// order.Status = "completed"
	// fmt.Println(order.Status)

	// // Create an instance of the struct without specifying the field names
	// order2 := Order{2, 200.0, "pending", time.Now(), "Mobile"}
	// fmt.Println(order2.Id)
	// fmt.Println(order2.Amount)
	// fmt.Println(order2.Status)
	// fmt.Println(order2.CreateDate)
	// fmt.Println(order2.Product)

	// // Call the receiver function
	// order.printOrder()

	// // Call the receiver function
	// order.updateStatus("shipped")
	// fmt.Println(order.Status)

	// order2.updateStatus("shipped")
	// fmt.Println(order2.Status)

	// Call the factory function
	order3 := newOrder(3, 300.0, "pending", time.Now(), "Tablet")
	order4 := newOrder(4, 400.0, "pending", time.Now(), "Desktop")
	order3.printOrder()
	order4.printOrder()
	order3.updateStatus("shipped") // This will not update the status of order4
	order3.printOrder()
	order4.printOrder()

	// Anonymous struct
	language := struct {
		Name string
		Version float64
	}{
		Name: "Go",
		Version: 1.15,
	}
	fmt.Println(language.Name)
	fmt.Println(language.Version)

}
