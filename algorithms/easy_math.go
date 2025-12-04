package algorithms

import "strconv"

func fizzBuzz(n int) []string {

	result := make([]string, n)
	for i := range n {
		result[i] = ""
		m := i + 1
		if m%3 == 0 {
			result[i] = "Fizz"
		}
		if m%5 == 0 {
			result[i] += "Buzz"
		}
		if len(result[i]) == 0 {
			result[i] = strconv.Itoa(m)
		}
	}
	return result
}
