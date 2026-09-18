package main

import "fmt"

func main() {
	fmt.Println(calculateTotal(30, 2))
	var val1,val2 = calculate(2,3)
	fmt.Println(sum(10, 20, 30))

	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(calculate2(3,4,add))

}

func calculateTotal(price, quantity float64) float64 {

	return price * quantity
}

func calculate(a, b float64) (float64, float64) {
	return a + b, a * b
}

func sum(numbers ...int) int{
 
	var totalSum = 0
	for _,val := range numbers {
        totalSum += val
	}

	return totalSum
}
func add(a, b int) int {
   return  a+b
}

func multiply(a, b int) int {
    return  a * b
}

func subtract(a, b int) int {
    return  a- b
}

func calculate2(a, b int, operation func(int, int) int) int {
   return operation(a,b)
}