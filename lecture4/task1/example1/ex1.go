package main

import (
	"fmt"
	"sync"
	//"runtime"
)

func main() {

	x := 0

	const num = 1000
	var wg sync.WaitGroup
	wg.Add(num * 3)

	for i := 0; i < num; i++ {
		go func() {
			curr := x
			//runtime.Gosched()
			curr++
			x = curr
			wg.Done()
		}()
	}

	for i := 0; i < num; i++ {
		go func() {
			curr := x
			//runtime.Gosched()
			curr = curr*2 + 1
			x = curr
			wg.Done()
		}()
	}

	for i := 0; i < num; i++ {
		go func() {
			curr := x
			//runtime.Gosched()
			curr = curr - 2
			x = curr
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", x)
}
