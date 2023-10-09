package main

import (
	"fmt"
	"sync"
)

func main() {
	x := make(chan int, 20)
	y := make(chan int, 20)
	result := make(chan int, 40)

	for i := 0; i < 20; i++ {
		x <- i
	}
	close(x)

	for j := 30; j < 50; j++ {
		y <- j
	}
	close(y)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for cur := range x {
			result <- cur
		}
	}()

	go func() {
		defer wg.Done()
		for cur := range y {
			result <- cur
		}
	}()

	wg.Wait()
	close(result)
	for i := 0; i < 40; i++ {
		fmt.Println(<-result)
	}
}
