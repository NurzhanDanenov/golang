package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters sync.Map
	var val int
	x := make(chan int, 100)
	counters.Store("habr", 25)
	v, ok := counters.Load("habr")
	if ok {
		val = v.(int)
	}
	v2, ok := counters.LoadOrStore("habr2", 13)
	fmt.Println(val, v2)
	counters.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true // if false, Range stops
	})
	wg := sync.WaitGroup{}
	//mu := sync.Mutex{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(id int) {
			//mu.Lock()
			counters.LoadOrStore(id, id)
			x <- id
			//mu.Unlock()
			wg.Done()
		}(i)
	}
	counters.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
		return true // Continue iterating
	})
	wg.Wait()
}
