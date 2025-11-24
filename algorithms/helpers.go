package algorithms

import "reflect"

// Flatten a 2D slice of strings
func flatten(matrix [][]string) []string {
	var flat []string
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

// Count occurrences of elements
func countElements(arr []string) map[string]int {
	counts := make(map[string]int)
	for _, val := range arr {
		counts[val]++
	}
	return counts
}

// Compare two matrices ignoring order
func matricesEqualIgnoreOrder(mat1, mat2 [][]string) bool {
	c1 := countElements(flatten(mat1))
	c2 := countElements(flatten(mat2))
	return reflect.DeepEqual(c1, c2)
}
