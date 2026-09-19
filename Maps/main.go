package main

import "fmt"

func main() {
	// mapped()
	// mapTask2()
	// mapTask3()
	// mapIterate()
	// lenCounting()
	nilMapOrEmty()
}

func mapped() {
	var mappedData = map[string]any{
		"name": "Torikul",
		"age":  19,
		"role": "Developer",
	}

	mappedData["age"] = 20
	mappedData["country"] = "Bangladesh"

	fmt.Println(mappedData["name"])
	fmt.Println(mappedData)
}

func pracOk() {
	users := map[string]string{
		"name": "Torikul",
		"role": "Developer",
	}

	if value, ok := users["name"]; ok {
		fmt.Println("name:", value)
	} else {
		fmt.Println("name:", "Key not found")
	}

	if value, ok := users["email"]; ok {
		fmt.Println("email:", value)
	} else {
		fmt.Println("email:", "Key not found")
	}

}

func mapDelete() {
	users := map[string]string{
		"name":    "Torikul",
		"role":    "Developer",
		"country": "Bangladesh",
	}

	delete(users, "role")

	if value, ok := users["role"]; ok {
		fmt.Println("role:", value)
	} else {
		fmt.Println("role:", "Key not found")
	}
	fmt.Println(users)
}

func mapIterate() {
	users := map[string]string{
		"name":    "Torikul",
		"role":    "Developer",
		"country": "Bangladesh",
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
}

func lenCounting() {
	users := map[string]string{
		"name":    "Torikul",
		"role":    "Developer",
		"country": "Bangladesh",
	}
	fmt.Println(len(users))

	for key := range users {
		fmt.Println("key", key)
	}
}

func nilMapOrEmty() {
	var users map[string]string
	users2 := make(map[string]string)

	fmt.Println(users)
	fmt.Println(users2)

	fmt.Println(users == nil)
	fmt.Println(users2 == nil)

	users2["name"] = "torikul"
	fmt.Println(users2)

	// users["name"] = "Torikul"
	// fmt.Println(users)

	// before i add key value in first map give panic error ase because this map is a nil map 

}
