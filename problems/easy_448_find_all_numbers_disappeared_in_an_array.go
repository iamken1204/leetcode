package problems

import "math"

// FindDisappearedNumbers represents
// Easy 448. Find All Numbers Disappeared in an Array
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
func FindDisappearedNumbers(nums []int) []int {
	var res []int
	var index int
	var key int

	for i := 0; i < len(nums); i++ {
		index = nums[i]
		key = int(math.Abs(float64(index))) - 1
		if nums[key] > 0 {
			nums[key] = -nums[key]
		}
	}

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			res = append(res, i)
		}
	}

	return res
}
