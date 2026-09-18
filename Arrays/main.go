package main

import "fmt"

func main() {
	// basicArray()
	// arrayCalculation()
	// arrayUpdateAndIndex()
	fixSize()
}

func basicArray() {
	var fiveInt = [5]int{1, 2, 3, 4, 5}

	for _, val := range fiveInt {
		fmt.Println(val)
	}

	fmt.Println("This is len", len(fiveInt))
}

func arrayCalculation() {
	var intArray = [5]int{10, 20, 30, 40, 50}
	var totalSum = 0
	var largestNumber = 0

	for _, val := range intArray {
		totalSum += val
		if largestNumber < val {
			largestNumber = val
		}
	}
	fmt.Println(totalSum)
	fmt.Println(largestNumber)
}

func arrayUpdateAndIndex() {
	var array = [5]int{10, 20, 30, 40, 50}
	lastIndex := len(array) - 1

	array[2] = 100
	array[0] = 500
	fmt.Println(array)
	fmt.Println("Arrays first element", array[0], "Arrays last element", array[lastIndex])
}

func fixSize() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr)
}
