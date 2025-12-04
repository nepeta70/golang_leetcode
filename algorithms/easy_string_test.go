package algorithms

import "testing"

func TestIsValidParentheses(t *testing.T) {

	input := "()[]{}"
	expected := true
	result := isValidParentheses(input)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
