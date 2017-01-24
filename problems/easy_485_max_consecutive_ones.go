package problems

import "math"

// FindMaxConsecutiveOnes represents
// Easy 485. Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones/
func FindMaxConsecutiveOnes(nums []int) int {
	var result float64
	var tmpCount float64
	var counting bool

	for _, b := range nums {
		if b == 1 {
			if !counting {
				counting = true
			}
			tmpCount++
		} else if counting {
			result = math.Max(tmpCount, result)
			tmpCount = 0
			counting = false
		}
	}
	result = math.Max(tmpCount, result)

	return int(result)
}
