package main

import (
	"fmt"
)

func Add[T int | float64 | string](a T, b T) T {
	return a + b
}

func main() {
	//   result := Add("sdf","sdf")

	//   fmt.Println(result)

	//   printValue("hello")
	//   printValue(444)
	// getFirst([]int{3, 4, 5})
	// getFirst([]string{"hello", "go"})
	fmt.Println(maxNum([]int{-5,-8}))
}

func printValue[T int | string](value T) {
	fmt.Println(value)
}

func getFirst[T int | string | float64](value []T) T {
	return value[0]
}

func maxNum[T int | float64](value []T) T {
	length := len(value)
	if length == 0 {
		return 0
	} else {
		var maxNumber T = value[0]
		for _, val := range value {
			if maxNumber < val {
				maxNumber = val
			}
		}
		return maxNumber
	}

}
