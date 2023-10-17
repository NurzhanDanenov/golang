package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		list1    *ListNode
		list2    *ListNode
		expected []int
	}{
		{
			list1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			list2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			list2: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			mergedList := mergeTwoLists(tc.list1, tc.list2)

			resultValues := []int{}
			current := mergedList
			for current != nil {
				resultValues = append(resultValues, current.Val)
				current = current.Next
			}

			if !reflect.DeepEqual(tc.expected, resultValues) {
				t.Errorf("Expected mergedValues %v, but got %v", tc.expected, resultValues)
			}
		})
	}
}
