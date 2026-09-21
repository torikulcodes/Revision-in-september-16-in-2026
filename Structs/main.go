package main

import "fmt"

type User struct {
	name        string
	age         int
	isDeveloper bool
}

type NewStruct struct {
	name string
}

func main() {
	u := User{"Torikul", 19, true}
	fmt.Println(u)
	fmt.Println(u.name)
	fmt.Println(u.age)
	u.age = 20
	fmt.Println(u.age)
	printUser(u)
	updateAge(&u)

	fmt.Println(u.age)

	namedStruct := User{
		name:        "Torikul",
		age:         19,
		isDeveloper: true,
	}
	fmt.Println(namedStruct)

	met := User{name: "Torikul"}
	fmt.Println("method",met.greet())
}

func printUser(u User) {
	fmt.Println(u.name)
	fmt.Println(u.age)
	fmt.Println(u.isDeveloper)
}

func updateAge(u *User) {
	u.age = u.age + 1
}

func (n User) greet() string{
	return fmt.Sprintf("Hello I am %s",n.name)
}
