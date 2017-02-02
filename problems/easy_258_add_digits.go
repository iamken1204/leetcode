package problems

// AddDigits represents
// Easy 258. Add Digits
// https://leetcode.com/problems/add-digits/
// 這個問題有公式... dr(n) = 1 + (n - 1) % 9
// https://en.wikipedia.org/wiki/Digital_root
func AddDigits(num int) int {
	return 1 + (num-1)%9
}
