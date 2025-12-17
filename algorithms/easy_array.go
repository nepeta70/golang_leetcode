package algorithms

import (
	"sort"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 1
	}

	count := 1
	lastVal := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] != lastVal {
			nums[count] = nums[i]
			count++
			lastVal = nums[i]
		}
	}

	return count
}

func maxProfit(prices []int) int {
	profit := 0
	for i := range len(prices) - 1 {
		var diff = prices[i+1] - prices[i]
		if diff > 0 {
			profit = profit + diff
		}
	}
	return profit
}

func rotate(nums []int, k int) {
	n := len(nums)
	if n < 2 {
		return
	}
	if k <= 0 {
		return
	}
	k = k % n

	arr := make([]int, k)

	for i := range k {
		arr[i] = nums[n-k+i]
	}

	for i := n - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}

	for i := range k {
		nums[i] = arr[i]
	}
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for i := range len(nums) {
		x := nums[i]
		_, exists := set[x]
		if exists {
			return true
		}
		set[x] = struct{}{}
	}
	return false
}

func singleNumber(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	sort.Ints(nums)

	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[n-1] != nums[n-2] {
		return nums[n-1]
	}
	for i := range len(nums) - 1 {
		if nums[i-1] != nums[i] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result []int
	k := 0
	for _, num1 := range nums1 {
		item1 := num1
		for j := k; j < len(nums2); j++ {
			item2 := nums2[j]
			if item1 == item2 {
				result = append(result, item1)
				k = j + 1
				break
			}
		}
	}
	return result
}

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		d := digits[i]
		if d < 9 {
			digits[i] = d + 1
			break
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func moveZeroes(nums []int) {

	if len(nums) <= 1 {
		return
	}
	zeroCount := 0
	step := 1

	for i := 1; i < len(nums); i++ {
		lastValue := nums[i-step]
		currentValue := nums[i]
		if lastValue == 0 {
			zeroCount++
			if currentValue != 0 {
				nums[i-step] = currentValue
				nums[i] = 0
			} else {
				step++
			}
		}
	}
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	for i := range n {
		column := make(map[byte]struct{})
		row := make(map[byte]struct{})
		square := make(map[byte]struct{})
		q0 := 3 * (i % 3)
		r0 := 3 * (i / 3)

		for j := range n {
			rowCell := board[i][j]
			if rowCell != '.' {
				_, exists := row[rowCell]
				if exists {
					return false
				}
				row[rowCell] = struct{}{}
			}

			colCell := board[j][i]
			if colCell != '.' {
				_, exists := column[colCell]
				if exists {
					return false
				}
				column[colCell] = struct{}{}
			}

			iq := q0 + j/3
			jq := r0 + j%3

			squareCell := board[iq][jq]
			if squareCell != '.' {
				_, exists := square[squareCell]
				if exists {
					return false
				}
				square[squareCell] = struct{}{}
			}
		}
	}
	return true
}

func rotateImage(matrix [][]int) {
	d0 := len(matrix) - 1
	d1 := d0
	d2 := d0

	// It'll rotate layer by layer from outside to inside
	for i := range d1 {
		for j := range d2 {
			col0 := i + j
			currentCell := matrix[i][col0]

			// new column index = initial row index
			// new row index = d0 - initial column index
			col1 := d0 - i
			targetCell := matrix[col0][col1]

			// Top to right
			matrix[col0][col1] = currentCell

			// new row index = previous column index
			// new column index = d0 - previous row index
			col2 := d0 - col0
			currentCell = matrix[col1][col2]

			// Right to bottom
			matrix[col1][col2] = targetCell

			// new column index = previous row index
			// new row index = d0 - previous column index
			col3 := d0 - col1
			targetCell = matrix[col2][col3]

			// Bottom to left
			matrix[col2][col3] = currentCell

			// new row index = previous column index = initial row index
			// new column index = d0 - previous row index = initial column index
			// Left to top
			matrix[i][col0] = targetCell
		}
		d1--
		d2 -= 2
	}
}
