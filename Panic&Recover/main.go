package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recover", r)
		}
		fmt.Println("After")
	}()
	checkAge(-3)
}

func checkAge(age int) {
	fmt.Println("Before")
	if age < 0 {
		panic("Invalid age")
	}
	fmt.Println("Valid age")

}
