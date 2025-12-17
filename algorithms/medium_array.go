package algorithms

import (
	"sort"
	"strings"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	if n == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			result = append(result, []int{nums[0], nums[1], nums[2]})
		}
		return result
	}

	// count := 1
	// lastVal := nums[0]
	// for i := 1; i < n; i++ {
	// 	if nums[i] != lastVal {
	// 		nums[count] = nums[i]
	// 		count++
	// 		lastVal = nums[i]
	// 	}
	// }

	// for i := range nums[:count-2] {
	// 	for j := range nums[i+1 : count-1] {
	// 		for k := range nums[j+1 : count] {
	// 			if i+j+k == 0 {
	// 				result = append(result, []int{i, j, k})
	// 			}
	// 		}
	// 	}
	// }
	lasti := nums[0]
	//lastj := nums[1]
	//lastk := nums[2]
	for i := 1; i < n-2; i++ {
		if nums[i] == lasti {
			continue
		}
		lasti = nums[i]
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if i != j && j != k && k != i {
					if nums[i]+nums[j]+nums[k] == 0 {
						result = append(result, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return result
}

func setZeroes(matrix [][]int) {

}

func groupAnagrams(strs []string) [][]string {

	n := len(strs)

	var result [][]string

	if n < 2 {
		result = append(result, []string{strs[0]})
		return result
	}
	dict := make(map[string][]string)

	for i := range n {
		item := strs[i]
		sorted := sortString(item)
		_, exists := dict[sorted]
		if !exists {
			dict[sorted] = []string{strs[i]}
		} else {
			dict[sorted] = append(dict[sorted], strs[i])
		}
	}

	for _, arr := range dict {
		result = append(result, arr)
	}

	return result
}

func lengthOfLongestSubstring(s string) int {

	n := len(s)
	if n < 2 {
		return n
	}

	longestLength := 0
	currentString := string(s[0])

	for i := 1; i < n; i++ {
		currentChar := string(s[i])
		containsChar := strings.Contains(currentString, currentChar)
		if !containsChar {
			currentString += currentChar
		} else {
			i := strings.LastIndex(currentString, currentChar)
			currentString = currentString[i+1:] + currentChar
		}
		currentLength := len(currentString)
		if currentLength > longestLength {
			longestLength = currentLength
		}
	}
	return longestLength
}
