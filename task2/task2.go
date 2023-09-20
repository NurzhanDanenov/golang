package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	input1 := []string{"flower", "flow", "flight"}
	result1 := longestCommonPrefix(input1)
	fmt.Println(result1)
	input2 := []string{"dog", "racecar", "car"}
	result2 := longestCommonPrefix(input2)
	fmt.Println(result2)
}
