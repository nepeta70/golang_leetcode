package algorithms

import "sort"

func topKFrequent(nums []int, k int) []int {
	n := len(nums)

	dict := make(map[int]int)
	var keys []int

	for i := 0; i < n; i++ {
		item := nums[i]
		_, exists := dict[item]
		if !exists {
			dict[item] = 1
			keys = append(keys, item)
		} else {
			dict[item] = dict[item] + 1
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return dict[keys[i]] > dict[keys[j]] // descending
	})

	return keys[:k]
}
