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
