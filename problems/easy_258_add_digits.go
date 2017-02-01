package problems

import (
	"log"
	"strconv"
	"strings"
)

// AddDigits represents
// Easy 258. Add Digits
// https://leetcode.com/problems/add-digits/
func AddDigits(num int) int {
	var res int
	var err error
	numS := strconv.Itoa(num)
	numSA := strings.Split(numS, "")
	if len(numSA) < 2 {
		res, err = strconv.Atoi(numSA[0])
		if err != nil {
			log.Fatalln(err)
		}
		return res
	}
	var next int
	for _, val := range numSA {
		tmp, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err)
		}
		next += tmp
	}
	return AddDigits(next)
}
