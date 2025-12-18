package algorithms

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	t.Parallel()
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	result := threeSum(nums)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSetZeroes(t *testing.T) {
	t.Parallel()
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	expected := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	setZeroes(matrix)

	if !reflect.DeepEqual(matrix, expected) {
		t.Errorf("Expected %v, got %v", expected, matrix)
	}
}

func TestGroupAnagrams(t *testing.T) {
	t.Parallel()
	strings := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	result := groupAnagrams(strings)

	if !matricesEqualIgnoreOrder(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	// 1. Define the test table
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "standard repeating pattern",
			input:    "abcabcbb",
			expected: 3, // "abc"
		},
		{
			name:     "repeating middle characters",
			input:    "pwwkew",
			expected: 3, // "wke"
		},
		{
			name:     "all identical characters",
			input:    "bbbbb",
			expected: 1, // "b"
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
	}

	// 2. Iterate over the table
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Maintain parallel execution from your original code

			// 3. Execute the function under test
			result := lengthOfLongestSubstring(tc.input)

			// 4. Assert the result
			if result != tc.expected {
				t.Errorf("For input %q: expected %d, got %d", tc.input, tc.expected, result)
			}
		})
	}
}
