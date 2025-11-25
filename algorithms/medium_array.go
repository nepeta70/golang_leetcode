package algorithms

import "strings"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	return [][]int{}
}

func setZeroes(matrix [][]int) {

}

func groupAnagrams(strs []string) [][]string {

	n := len(strs)

	var result [][]string

	if len(strs) < 2 {
		result = append(result, []string{strs[0]})
		return result
	}
	dict := make(map[string][]string)

	for i := 0; i < n; i++ {
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
