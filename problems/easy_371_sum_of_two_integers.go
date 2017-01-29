package problems

// GetSum represents
// Easy 371. Sum of Two Integers
// https://leetcode.com/problems/sum-of-two-integers/
func GetSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		a, b = a^b, (a&b)<<1
	}

	return a
}
