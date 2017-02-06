package problems

// MoveZeroes represents
// 1283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
func MoveZeroes(nums []int) {
	var idx int
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}
	for idx < len(nums) {
		nums[idx] = 0
		idx++
	}
}
