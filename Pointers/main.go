package main

import "fmt"

func main() {
	// pracPointers()
	age := 19
	a := 10
	b := 20
	fmt.Println(increment(&age))
	// fmt.Println(age)
	pointerSwap(&a, &b)



}

func pracPointers() {
	age := 19

	var ptr *int = &age

	var ptrToValue = *ptr

	fmt.Println(ptrToValue)

	*ptr = 20

	fmt.Println(age)
}

func increment(age *int) int {
	*age = *age + 1

	return *age
}

func pointerSwap(firstPointer, secondPointer *int) {
	storeValue := *firstPointer
	*firstPointer = *secondPointer
	*secondPointer = storeValue
	fmt.Println(*firstPointer)
	fmt.Println(*secondPointer)
}
