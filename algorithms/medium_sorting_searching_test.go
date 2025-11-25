package algorithms

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {

	input := []int{1, 1, 1, 2, 2, 3}

	expected := []int{1, 2}
	k := 2
	result := topKFrequent(input, k)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMerge(t *testing.T) {

	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	expected := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}

	result := merge(intervals)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
