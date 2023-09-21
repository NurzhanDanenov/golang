package main

import "fmt"

func main() {
	num1 := 3
	ex1 := intToRoman(num1)
	fmt.Println(ex1)

	num2 := 58
	ex2 := intToRoman(num2)
	fmt.Println(ex2)

	num3 := 1994
	ex3 := intToRoman(num3)
	fmt.Println(ex3)
}

func intToRoman(num int) string {
	romanDigits := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	digits := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for i := 0; i < len(digits); i++ {
		for num >= digits[i] {
			result += romanDigits[i]
			num -= digits[i]
		}
	}

	return result
}
