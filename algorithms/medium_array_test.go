package algorithms

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
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
	strings := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	result := groupAnagrams(strings)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
