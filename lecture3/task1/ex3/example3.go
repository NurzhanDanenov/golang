package main

func main() {
	x := 1
	y := make(chan int, 3)
	for i := 0; i < 4; i++ {
		y <- x + i
	}
}
