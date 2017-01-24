package problems

import "strings"
import "bytes"

// ReverseString represents
// Easy 344 Reverse String.
// https://leetcode.com/problems/reverse-string/
func ReverseString(s string) string {

	var res bytes.Buffer
	var x string

	arr := strings.Split(s, "")

	for len(arr) > 0 {
		x, arr = arr[len(arr)-1], arr[:len(arr)-1]
		res.WriteString(x)
	}

	return res.String()

}
