package problems

import "strconv"

// FizzBuzz represents
// Easy 412. Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
func FizzBuzz(n int) []string {
	var m3 bool
	var m5 bool
	var res []string
	var attach string

	for i := 1; i <= n; i++ {
		m3 = (i % 3) == 0
		m5 = (i % 5) == 0
		if m3 && m5 {
			attach = "FizzBuzz"
		} else if m3 {
			attach = "Fizz"
		} else if m5 {
			attach = "Buzz"
		} else {
			attach = strconv.Itoa(i)
		}
		res = append(res, attach)
	}

	return res
}
