package main

import "fmt"

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	res1 := twoSum(nums1, target1)
	fmt.Println(res1)

	nums2 := []int{3, 2, 4}
	target2 := 6
	res2 := twoSum(nums2, target2)
	fmt.Println(res2)

	nums3 := []int{3, 3}
	target3 := 6
	res3 := twoSum(nums3, target3)
	fmt.Println(res3)
}

func twoSum(nums []int, target int) []int {
	x := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		b := target - nums[i]
		if idx, ok := x[b]; ok {
			return []int{i, idx}
		} else {
			x[nums[i]] = i
		}
	}
	return []int{}
}
