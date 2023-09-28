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
			temp := x
			//runtime.Gosched()
			temp++
			x = temp
			wg.Done()
		}()
	}

	for i := 0; i < num; i++ {
		go func() {
			temp := x
			//runtime.Gosched()
			temp = temp*2 + 1
			x = temp
			wg.Done()
		}()
	}

	for i := 0; i < num; i++ {
		go func() {
			temp := x
			//runtime.Gosched()
			temp = temp - 2
			x = temp
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", x)
}
