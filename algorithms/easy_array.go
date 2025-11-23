package algorithms

import (
	"sort"
)

func removeDuplicates(nums []int) int {
	return 0
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
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

	for i := 0; i < k; i++ {
		nums[i] = arr[i]
	}
}

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func singleNumber(nums []int) int {
	return 0
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result []int
	k := 0
	for i := 0; i < len(nums1); i++ {
		item1 := nums1[i]
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
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
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
	d := len(matrix)
	y0 := 0
	d0 := d - 1
	d1 := d0
	d2 := d1

	for i := y0; i < d1; i++ {
		for j := 0; j < d2; j++ {
			row0 := i
			col0 := i + j
			previousCell := matrix[row0][col0]

			// new column index = initial row index
			// new row index = d0 - initial column index
			col1 := d0 - row0
			nextCell := matrix[col0][col1]

			// Top to right
			matrix[col0][col1] = previousCell

			// new row index = previous column index
			// new column index = d0 - previous row index
			col2 := d0 - col0
			previousCell = matrix[col1][col2]

			// Right to bottom
			matrix[col1][col2] = nextCell

			// new column index = previous row index
			// new row index = d0 - previous column index
			col3 := d0 - col1
			nextCell = matrix[col2][col3]

			// Bottom to left
			matrix[col2][col3] = previousCell

			// new row index = previous column index = initial row index
			// new column index = d0 - previous row index = initial column index
			// Left to top
			matrix[row0][col0] = nextCell
		}
		y0++
		d1--
		d2 -= 2
	}
}
