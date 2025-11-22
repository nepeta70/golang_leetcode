package algorithms

import (
	"reflect"
	"testing"
)

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
