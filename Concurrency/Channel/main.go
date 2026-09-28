package main

import "fmt"

func main() {
	// taskOne()
	// taskTwo()
	// taskThree()
	taskFour()
}

func taskOne() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()

	fmt.Println(<-ch)
}

func taskTwo() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
	}()

	// a := <-ch
	// b := <-ch
	// c := <-ch
	// fmt.Println(a, b, c)

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-ch)
	// }

}

func taskThree() {
	ch := make(chan int, 1)

	ch <- 100

	fmt.Println(<-ch)
}

func taskFour() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch)
}
