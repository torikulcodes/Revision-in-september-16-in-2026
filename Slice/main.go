package main

import "fmt"

func main() {
	addItems()
	slicing()
	copySlice()
	lenCap()
}

func addItems() {
	numbers := []int{10, 20, 30}
	numbers = append(numbers, 40)
	numbers = append(numbers, 50, 60)

	fmt.Println(numbers)
}

func slicing() {
	numbers := []int{10, 20, 30, 40, 50, 60}

	slice1 := numbers[1:4]
	slice2 := numbers[len(numbers)-3:]
	slice3 := numbers[0:3]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func copySlice() {
	source := []int{10, 20, 30, 40, 50}
	destination := make([]int, len(source))

	copy(destination, source)
	fmt.Println(destination)

	source[0] = 100
	fmt.Println(source)
	fmt.Println(destination)
}

func lenCap() {
	numbers := make([]int, 3, 5)
	fmt.Println(numbers)

	fmt.Println("length", len(numbers))
	fmt.Println("cap", cap(numbers))

	numbers = append(numbers, 6, 7)

	fmt.Println("length", len(numbers))
	fmt.Println("cap", cap(numbers))
}
