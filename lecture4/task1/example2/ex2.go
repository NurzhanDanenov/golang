package main

import (
	"fmt"
	//"runtime"
	"sync"
)

func main() {

	x := 1000
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		curr := x
		//runtime.Gosched()
		curr = curr + 100
		x = curr
	}()
	go func() {
		defer wg.Done()
		curr := x
		//runtime.Gosched()
		curr = curr / 2
		x = curr
	}()
	go func() {
		defer wg.Done()
		curr := x
		//runtime.Gosched()
		curr = curr * 2
		x = curr
	}()
	go func() {
		defer wg.Done()
		curr := x
		//runtime.Gosched()
		curr = curr - 100
		x = curr
	}()
	wg.Wait()
	fmt.Println("count:", x)
}
