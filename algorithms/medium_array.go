package algorithms

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
