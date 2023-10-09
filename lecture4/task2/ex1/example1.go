package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]int)
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(id int) {
			mu.Lock()
			myMap[id] = id
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		mu.Lock()
		val, ok := myMap[i]
		mu.Unlock()
		if ok {
			fmt.Printf("Key: %v, Value: %v\n", i, val)
		}
	}
}
