package algorithms

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper to find a specific node by value so we can pass it to deleteNode
func findNode(head *ListNode, val int) *ListNode {
	curr := head
	for curr != nil {
		if curr.Val == val {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		valToDelete   int
		expectedSlice []int
	}{
		{
			name:          "Delete 5 from 4,5,1,9",
			input:         []int{4, 5, 1, 9},
			valToDelete:   5,
			expectedSlice: []int{4, 1, 9},
		},
		{
			name:          "Delete 1 from 4,5,1,9",
			input:         []int{4, 5, 1, 9},
			valToDelete:   1,
			expectedSlice: []int{4, 5, 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			head := sliceToList(tc.input)
			targetNode := findNode(head, tc.valToDelete)
			expectedHead := sliceToList(tc.expectedSlice)

			// Act
			deleteNode(targetNode)

			// Assert
			if !reflect.DeepEqual(head, expectedHead) {
				t.Errorf("Resulting list mismatch. Expected %v, but got %v", tc.expectedSlice, head)
			}
		})
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "Remove 2nd from end of 5 elements",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Remove only element",
			input:    []int{1},
			n:        1,
			expected: nil,
		},
		{
			name:     "Remove last element of two",
			input:    []int{1, 2},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: Convert slices to Linked Lists
			head := sliceToList(tc.input)
			expectedHead := sliceToList(tc.expected)

			// Act
			result := removeNthFromEnd(head, tc.n)

			// Assert
			// reflect.DeepEqual works well for nested pointer structures like linked lists
			if !reflect.DeepEqual(result, expectedHead) {
				t.Errorf("For n=%d, expected %v, but got %v", tc.n, expectedHead, result)
			}
		})
	}
}
