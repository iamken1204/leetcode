package problems

// TwoSum2 represents
// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSum2(numbers []int, target int) []int {
	var res []int
	var i int
	var done bool
	maxI := len(numbers) - 2
	for i <= maxI && numbers[i] <= target {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				res = []int{i + 1, j + 1}
				done = true
				break
			}
		}
		if done {
			break
		}
		i++
	}
	return res
}
