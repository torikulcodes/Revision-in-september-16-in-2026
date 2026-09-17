package main

import "fmt"

func main() {
	fmt.Println(userStatus(33, "sd"))
	numbers := []int{5,8,-3,0}
	analyzeNumbers(numbers)
}

func userStatus(age int, status string) string {
	var ageGroup string

	if age < 13 {
		ageGroup = "child"
	} else if age >= 13 && age < 18 {
		ageGroup = "Teenager"
	} else {
		ageGroup = "Adult"
	}

	switch status {
	case "active":
		return ageGroup + " - User is active"
	case "inactive":
		return ageGroup + " - User is inactive"
	default:
		return ageGroup + " - Unknown status"
	}
}

func analyzeNumbers(numbers []int) {

	for _, val := range numbers {
		if val == 0 {
			continue
		} else if val%2 == 0 && val >= 1 {
			fmt.Println(val, "-> Positive, Even")
		} else if val%2 == 0 && val < 0 {
			fmt.Println(val, "-> Negative, Even")
		} else if val%2 != 0 && val >= 1 {
			fmt.Println(val, "-> Positive, Odd")
		} else if val%2 != 0 && val < 0 {
			fmt.Println(val, "-> Negative, Odd")
		}
	}

}
