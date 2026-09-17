package main

import "fmt"

func main() {
	taskOne()
	taskTwo()
	taskThree()
}

func taskOne() {
	var name = "Torikul"
	var age = 19
	var salary = 14000
	var isRemote = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(salary)
	fmt.Println(isRemote)
}

func taskTwo() {
	var name string = "Torikul islam"
	var age int = 19
	var height float64 = 5.6
	var grade = 'A'
	var isStudent bool = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(grade)
	fmt.Println(isStudent)
}

func taskThree() {
	var name = "torikul"
	var monthlySalary = 28000
	const tax = 10

	fmt.Println(name)
	fmt.Println(monthlySalary)
	fmt.Println("tax", monthlySalary / 100 * tax)
}
