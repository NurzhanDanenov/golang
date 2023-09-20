package main

import (
	"fmt"
	"sort"
)

type IntegerSlice struct {
	data []int
}

func (is *IntegerSlice) Add(values ...int) {
	is.data = append(is.data, values...)
}

func (is *IntegerSlice) Sort() {
	sort.Ints(is.data)
}

func (is *IntegerSlice) Print() {
	fmt.Println(is.data)
}

func main() {
	intSlice := IntegerSlice{}

	intSlice.Add(5, 3, 9, 1, 7)

	fmt.Println("Слайс до сортировки:")
	intSlice.Print()

	intSlice.Sort()

	fmt.Println("Слайс после сортировки:")
	intSlice.Print()
}
