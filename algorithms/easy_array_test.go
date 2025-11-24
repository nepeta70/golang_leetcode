package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	input := []int{1, 1, 2}
	result := removeDuplicates(input)

	expected := 2
	output := []int{1, 2}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("input: %v\n", input)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
		for i := 0; i < result; i++ {
			if input[i] != output[i] {
				t.Errorf("Output different, expected %v, real %v", output, input)
			}
		}
	}
}

func TestRemoveDuplicates2(t *testing.T) {

	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(input)

	expected := 5
	output := []int{0, 1, 2, 3, 4}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("input: %v\n", input)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
		for i := 0; i < result; i++ {
			if input[i] != output[i] {
				t.Errorf("Output different, expected %v, real %v", output, input)
			}
		}
	}
}

func TestMaxProfit(t *testing.T) {

	input := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(input)

	if result != 7 {
		t.Errorf("Expected 7, got %d", result)
	}
}

func TestRotate(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)

	expected := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestContainsDuplicate(t *testing.T) {

	input := []int{1, 2, 3, 1}
	result := containsDuplicate(input)

	if result != true {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestSingleNumber(t *testing.T) {

	input := []int{4, 1, 2, 1, 2}
	result := singleNumber(input)

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestIntersect(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2, 2}
	result := intersect(nums1, nums2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPlusOne(t *testing.T) {
	input := []int{1, 2, 3}
	result := plusOne(input)

	expected := []int{1, 2, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestMoveZeroes(t *testing.T) {
	input := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(input, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result := isValidSudoku(board)
	if result != true {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestRotateImage(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	rotateImage(matrix)

	if !reflect.DeepEqual(matrix, expected) {
		t.Errorf("Expected %v, got %v", expected, matrix)
	}
}
