package main

import "fmt"

func main() {
	defer fmt.Println("Close resource")
	defer fmt.Println("Do some work")
	defer fmt.Println("Open resource")

}
