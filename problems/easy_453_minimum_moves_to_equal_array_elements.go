package problems

// MinMoves represents
// Easy 453. Minimum Moves to Equal Array Elements
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/
//
// As for `A Step`:
// Add 1 to all numbers except nums.Max === Minus 1 to a single number which is not nums.Min
func MinMoves(nums []int) int {
	var res int
	min := nums[0]

	// get the minimum number in nums
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	// count moves
	for _, num := range nums {
		res += num - min
	}

	return res
}
