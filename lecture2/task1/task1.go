package main

import "fmt"

func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumeral := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			romanNumeral += symbols[i]
			num -= values[i]
		}
	}

	return romanNumeral
}

func main() {
	integer := 3549
	roman := intToRoman(integer)
	fmt.Printf("The Roman numeral representation of %d is %s\n", integer, roman)
}
