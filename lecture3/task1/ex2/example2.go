package main

import (
	"fmt"
)

func main() {
	x := make(chan int)

	go channel_reading(x)

	for {
		msg := <-x
		fmt.Println(msg)
	}
}

func channel_reading(x chan int) {
	x <- 890
	x <- 350
}
