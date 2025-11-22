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
