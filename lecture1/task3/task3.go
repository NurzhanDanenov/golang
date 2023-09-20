package main

import "fmt"

func main() {
	slc1 := []int{2, 6, 7, 6}
	slc2 := []int{6, 6, 2, 7}
	res := equalsWithoutOrder(slc1, slc2)
	fmt.Println(res)

	slc1 = []int{2, 6, 7}
	slc2 = []int{2, 6, 7}
	res = equalsWithOrder(slc1, slc2)
	fmt.Println(res)
}

func equalsWithOrder(slc1 []int, slc2 []int) bool {
	if len(slc1) == len(slc2) {
		for i := 0; i < len(slc1); i++ {
			if slc1[i] != slc2[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func equalsWithoutOrder(slc1 []int, slc2 []int) bool {
	if len(slc1) != len(slc2) {
		return false
	}
	a := make(map[int]int, 0)
	b := make(map[int]int, 0)

	for _, aVal := range slc1 {
		a[aVal]++
	}

	for _, bVal := range slc2 {
		b[bVal]++
	}

	for aKey, aVal := range a {
		if b[aKey] != aVal {
			return false
		}
	}
	return true
}
