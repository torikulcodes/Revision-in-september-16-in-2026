package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Paid
	Shipped
	Delivered
	Cancelled
)

func main() {

	fmt.Println(statusMessage(Cancelled))
}

func statusMessage(status Status) string {
	switch status {
	case Pending:
		return "Waiting for payment"
	case Paid:
		return "Payment successful"
	case Shipped:
		return "Order shipped"
	case Delivered:
		return "Order delivered"
	case Cancelled:
		return "Order cancelled"
	default:
		return "Invalid status"
	}
}
