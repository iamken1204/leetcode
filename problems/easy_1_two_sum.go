package problems

// TwoSum represents
// Easy 1. Two Sum
// https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {

	res := make([]int, 2, 2)

	var L int

	maxI := len(nums) - 1
	maxJ := len(nums)
	flagI := 0
	flagJ := flagI + 1

	for i := flagI; i < maxI; i++ {
		L = nums[i]

		for j := flagJ; j < maxJ; j++ {
			if (L + nums[j]) == target {
				res[0] = i
				res[1] = j
				return res
			}
		}

		flagJ++
	}

	return res
}
