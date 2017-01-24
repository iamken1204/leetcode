package problems

import (
	"bytes"
	"fmt"
	"strconv"
)

// FindComplement represents
// Easy 476. Number Complement
// https://leetcode.com/problems/number-complement/
func FindComplement(num int) int {
	// Transform given number to binary (string).
	// 5 --> "101"
	binStr := strconv.FormatInt(int64(num), 2)

	// Announce a bytes.Buffer,
	// loop binary (string), reverse "0" to "1" and vice versa,
	// then write into bytes.Buffer.
	var reversed bytes.Buffer
	for i := 0; i < len(binStr); i++ {
		check := fmt.Sprintf("%c", binStr[i])
		if check == "1" {
			reversed.WriteString("0")
		} else {
			reversed.WriteString("1")
		}
	}

	// Transform binary (string) to int32.
	res, _ := strconv.ParseInt(reversed.String(), 2, 64)

	return int(res)
}
