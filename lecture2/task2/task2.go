package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	mergedList := mergeTwoLists(list1, list2)
	fmt.Println(mergedList)

	list1 = &ListNode{}
	list2 = &ListNode{}
	mergedList = mergeTwoLists(list1, list2)
	fmt.Println(mergedList)

	list1 = &ListNode{}
	list2 = &ListNode{Val: 0}
	mergedList = mergeTwoLists(list1, list2)
	fmt.Println(mergedList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedValues := []int{}
	for list1 != nil {
		mergedValues = append(mergedValues, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		mergedValues = append(mergedValues, list2.Val)
		list2 = list2.Next
	}

	sort.Ints(mergedValues)

	var mergedList *ListNode
	current := mergedList
	for _, val := range mergedValues {
		node := &ListNode{Val: val}
		if current == nil {
			current = node
			mergedList = current
		} else {
			current.Next = node
			current = current.Next
		}
	}

	return mergedList
}
