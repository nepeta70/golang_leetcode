package algorithms

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	input := 15
	result := fizzBuzz(input)
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
