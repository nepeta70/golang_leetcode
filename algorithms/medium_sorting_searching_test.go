package algorithms

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	t.Parallel()
	input := []int{1, 1, 1, 2, 2, 3}

	expected := []int{1, 2}
	k := 2
	result := topKFrequent(input, k)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMerge(t *testing.T) {
	// 1. Define the table of test cases
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "Standard overlapping intervals",
			input: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},
			expected: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
		{
			name: "Fully contained interval",
			input: [][]int{
				{1, 4}, {2, 3},
			},
			expected: [][]int{
				{1, 4},
			},
		},
		{
			name: "No overlaps",
			input: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
			expected: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
		},
		{
			name: "Touch at boundaries",
			input: [][]int{
				{1, 4}, {4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
	}

	// 2. Iterate through the table
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Runs subtests in parallel

			// 3. Execute the function
			result := merge(tc.input)

			// 4. Validate results
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test '%s' failed:\nExpected: %v\nGot:      %v",
					tc.name, tc.expected, result)
			}
		})
	}
}
