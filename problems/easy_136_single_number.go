package problems

// SingleNumber represents
// Easy 136. Single Number
// https://leetcode.com/problems/single-number/
func SingleNumber(nums []int) int {

	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]

}
