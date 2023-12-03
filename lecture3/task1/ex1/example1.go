package main

import (
	"fmt"
	"time"
)

func main() {
	x := make(chan string)

	go channel_string("Nurzhik is the best", x)

	for {
		msg := <-x
		fmt.Println(msg)
	}
}

func channel_string(smt string, c chan string) {
	for i := 0; i < 4; i++ {
		c <- smt
		time.Sleep(time.Microsecond * 500)
	}
}
