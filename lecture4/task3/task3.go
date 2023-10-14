package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]int)
	var mu sync.RWMutex
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func(id int) {
			mu.Lock()
			myMap[id] = id
			mu.Unlock()
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(id int) {
			mu.RLock()
			val, ok := myMap[i]
			mu.RUnlock()
			if ok {
				fmt.Printf("Key: %v, Value: %v\n", i, val)
			}
			wg.Done()
		}(i)
	}

	go func() {
		mu.RLock()
		for i := range myMap {
			val, ok := myMap[i]
			if ok {
				fmt.Printf("Key: %v, Value: %v\n", i, val)
			}
		}
		mu.RUnlock()
		wg.Done()
	}()
	wg.Wait()
}
