package problems

import "sort"

// FindContentChildren represents
// Easy 455. Assign Cookies
// https://leetcode.com/problems/assign-cookies/
//
// g => children
// s => cookies
func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var idxG int
	var idxS int
	for idxG < len(g) && idxS < len(s) {
		if s[idxS] >= g[idxG] {
			idxG++
			idxS++
		} else {
			idxS++
		}
	}
	return idxG
}
