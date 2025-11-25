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

func merge(intervals [][]int) [][]int {

	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastInterval := intervals[0]
	var result [][]int
	for i := 1; i <= n; i++ {
		if i == n {
			result = append(result, lastInterval)
		} else {
			interval := intervals[i]
			start0, end0 := lastInterval[0], lastInterval[1]
			start1, end1 := interval[0], interval[1]
			if start1 <= end0 {
				latestEnd := end0
				if end1 > end0 {
					latestEnd = end1
				}
				lastInterval = []int{start0, latestEnd}
			} else {
				result = append(result, lastInterval)
				lastInterval = []int{start1, end1}
			}
		}
	}
	return result
}
