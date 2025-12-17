package algorithms

import "strconv"

func fizzBuzz(n int) []string {

	result := make([]string, n)
	for i := range n {
		m := i + 1
		switch {
		case m%3 == 0 && m%5 == 0:
			result[i] = "FizzBuzz"
		case m%3 == 0:
			result[i] = "Fizz"
		case m%5 == 0:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(m)
		}
	}
	return result
}
