package main

import "fmt"

type OrderStatus string

const (
	Pending OrderStatus = "Pending"
	Processing OrderStatus = "Processing"
	Shipped OrderStatus = "Shipped"
	Delivered OrderStatus = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status is: ", status)
}

func main() {
	changeOrderStatus(Pending)
	changeOrderStatus(Processing)
	changeOrderStatus(Shipped)
	changeOrderStatus(Delivered)
}
