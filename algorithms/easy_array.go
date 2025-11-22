package algorithms

func removeDuplicates(nums []int) int {
	return 0
}

func maxProfit(prices []int) int {
	var profit int = 0
	for i := 0; i < len(prices)-1; i++ {
		var diff = prices[i+1] - prices[i]
		if diff > 0 {
			profit = profit + diff
		}
	}
	return profit
}

func rotate(nums []int, k int) {
	var n = len(nums)
	if n < 2 {
		return
	}
	if k <= 0 {
		return
	}
	k = k % n

	var arr = make([]int, k)

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
	return false
}

func singleNumber(nums []int) int {
	return 0
}

func intersect(nums1 []int, nums2 []int) []int {
	return []int{}
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

}

func twoSum(nums []int, target int) []int {
	return []int{}
}

func isValidSudoku(board [][]byte) bool {
	return false
}

func rotateImage(matrix [][]int) {

}
