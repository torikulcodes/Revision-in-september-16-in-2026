package main

import (
	"fmt"
	"time"
)

func main() {
	// taskOne()
	// taskTwo()
	// taskThree()
	// taskFour()
	// taskFive()

	// ch := make(chan string)
	// for i := 1; i < 4; i++ {
	// 	go taskSix(i, ch)

	// }

	// for i := 1; i < 4; i++ {
	// 	fmt.Println(<-ch)
	// }

	// taskSeven()
	taskEight()

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

func taskFive() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		// এখানে কী করতে হবে?
		close(ch)

	}()

	// এখানে এমনভাবে receive করো
	// যাতে 10, 20, 30 তিনটিই পাওয়া যায়
	// এবং channel-এর শেষ বুঝে loop থেমে যায়।
	for value := range ch {
		fmt.Println(value)
	}
}

func taskSix(id int, ch chan string) {
	ch <- fmt.Sprintf("Worker %d finished", id)
}

func taskSeven() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch2 <- "Message from channel 2"
	}()
	go func() {
		ch1 <- "Message from channel 1"
	}()
	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	}

	// routineOne := 0
	// routineTwo := 0
	// for i := 0; i < 100; i++ {
	// 	ch1 := make(chan string)
	// 	ch2 := make(chan string)

	// 	go func() {
	// 		ch2 <- "Message from channel 2"
	// 	}()
	// 	go func() {
	// 		ch1 <- "Message from channel 1"
	// 	}()
	// 	select {
	// 	case res := <-ch1:
	// 		fmt.Println(res)
	// 		routineOne++
	// 	case res := <-ch2:
	// 		fmt.Println(res)
	// 		routineTwo++
	// 	}
	// }

	// fmt.Println("routine one", routineOne, "routine two", routineTwo)

}

func taskEight() {
	ch := make(chan string)

	// এমন একটি goroutine তৈরি করো
	// যেটা 3 second পরে ch-এ
	// "Task completed" পাঠাবে

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "completed"
	}()


	// কিন্তু main goroutine সর্বোচ্চ 1 second অপেক্ষা করবে।

	// 1 second-এর মধ্যে message এলে:
	// Task completed

	// না এলে:
	// Timeout
	start := time.Now()
	select {
	case result := <-ch:
		fmt.Println(result)

	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}

	fmt.Println("execution time ",time.Since(start))
}
