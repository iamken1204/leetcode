package problems

// CanConstruct represents
// Easy 383. Ransom Note
// https://leetcode.com/problems/ransom-note/
func CanConstruct(ransomNote string, magazine string) bool {
	check := make(map[rune]int, len(ransomNote))
	for _, r := range magazine {
		check[r]++
	}
	for _, m := range ransomNote {
		if _, ok := check[m]; ok {
			check[m]--
			if check[m] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
