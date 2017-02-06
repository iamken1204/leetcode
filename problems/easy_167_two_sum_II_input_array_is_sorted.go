package problems

// TwoSum2 represents
// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSum2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for numbers[left]+numbers[right] != target {
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}
