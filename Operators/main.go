package main

import "fmt"

func main() {
	arithmeticOperators(2, 3)
	checkNumber(5)
	assignmentOperators()
	counter()
}

func arithmeticOperators(a, b int) {
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}

func checkNumber(n int) {
	var isSmall = n < 18
	var isBig = n > 18
	var isSimilar = n == 18
	var isSimilarOrBig = n >= 18
	var isSimilarOrSmall = n <= 18
	var notEqual = n != 18
	var useAnd = n > 10 && n < 20
	var useOR = n == 10 || n == 5
	isRaining := true
	takeUmbrella := !isRaining

	fmt.Println(isSmall)
	fmt.Println(isBig)
	fmt.Println(isSimilar)
	fmt.Println(isSimilarOrBig)
	fmt.Println(isSimilarOrSmall)
	fmt.Println(notEqual)
	fmt.Println(useAnd)
	fmt.Println(useOR)
	fmt.Println(takeUmbrella)
}

func assignmentOperators() {
	score := 10
	fmt.Println(score)
	score += 10
	fmt.Println(score)
	score -= 10
	fmt.Println(score)
	score *= 10
	fmt.Println(score)
	score /= 10
	fmt.Println(score)
	score %= 10
	fmt.Println(score)
}

func counter() {
	count := 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)

}
