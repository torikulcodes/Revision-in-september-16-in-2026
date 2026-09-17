package main

import "fmt"

func main() {
	// fmt.Println(float64(44))
	dataTypes()
	typeConversion()
	runeByte()
}

func dataTypes() {
	var name string = "Torikul"
	var age int = 19
	var height float64 = 5.6
	var isDeveloper bool = true
	var grade rune = '😊'

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(isDeveloper)
	fmt.Println(grade)
}

func typeConversion() {
	var age = 19
	var height = 5.6
	ageFloat := float64(age)

	heightInt := int(height)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(ageFloat)
	fmt.Println(heightInt)

}

func runeByte() {
	var letter rune = 'A'
	var symbol byte = 'B'
	fmt.Println(letter)
	fmt.Println(symbol)
	fmt.Println(letter + 1)
	fmt.Println(symbol + 1)
}
