package main

import "fmt"

func main() {
a := nextInt()
b := nextInt()

fmt.Println("this is a",a())
fmt.Println("this is a",a())
fmt.Println("this is a",a())
fmt.Println("this is a",b())
fmt.Println("this is a",b())
fmt.Println("this is a",b())



	// anonymousFunction in example
	// 1. one way 
	anonymousFunc := func (str string)string  {
		return str
	}
	anonymousFunc("hello")

	// two way iife
	func (str string)  {
		fmt.Println(str)
	}("hello")

	// task for chatgpt

	func (a,b int)int  {
		return a + b
	}(3,4)

	multiply := func (a,b int)int  {
		return a*b
	}
	multiply(2,3)
}

// anonymous function in example

func nextInt() func() int {
	value := 0

	return func() int {
		value++
		return value
	}
}
