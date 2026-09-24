package main

import (
	"errors"
	"fmt"
)

func main() {
	// result, error := divide(20, 0)
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// fmt.Println(result)

	// age,err := getUserAge(-9)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(age)
	// }

	// age2,err := getUserAge(5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(age2)
	// }

	fmt.Println(checkAge(-4))

	fmt.Println(checkAge(4))


}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func getUserAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("Age cannot be negative")
	} else if age >= 0 {
		return "Valid age", nil
	}
	return "nothing", nil
}

func checkAge(age int)error{
	if age < 0 {
		return fmt.Errorf("age must be grater then 0 your age is %d",age)
	}else if age >= 0 {
		return nil
	}
	return fmt.Errorf("Error")
}