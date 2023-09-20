package main

import (
	"fmt"
	"sort"
)

type Slice struct {
	data []int
}

func (is *Slice) Add(values ...int) {
	is.data = append(is.data, values...)
}

func (is *Slice) Sort() {
	sort.Ints(is.data)
}

func (is *Slice) Print() {
	fmt.Println(is.data)
}

func main() {
	intSlice := Slice{}

	intSlice.Add(5, 3, 9, 1, 7)

	fmt.Println("Слайс до сортировки:")
	intSlice.Print()

	intSlice.Sort()

	fmt.Println("Слайс после сортировки:")
	intSlice.Print()
}
