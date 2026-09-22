package main

import "fmt"

type PaymentMethod interface {
	Pay() string
}

type Bkash struct {
}
type Nagad struct{

}
func main() {

	// var p PaymentMethod
	// p = Bkash{}
	// fmt.Println(p.Pay())
	// p = Nagad{}
	// fmt.Println(p.Pay())
	processPayment(Bkash{})
	processPayment(Nagad{})
}

func processPayment(p PaymentMethod){
	fmt.Println(p.Pay())
}

func (b Bkash) Pay() string {
	return "Bkash payment successful"
}
func(n Nagad) Pay()string {
  return "Nagad payment successful"
}