package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printLetters() {

	for _, letter := range []string{"A", "B", "C", "D", "E"} {
		fmt.Println(letter)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// defer wg.Done()
	wg.Go(printNumbers)
	wg.Go(printLetters)
	wg.Wait()

}
