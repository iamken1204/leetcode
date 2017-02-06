package problems

// FindContentChildren represents
// Easy 455. Assign Cookies
// https://leetcode.com/problems/assign-cookies/
func FindContentChildren(g []int, s []int) int {
	var content int
	var cookie int
	var res int
	var resTmp int
	var chosenIdx int
	remainSteps := len(s)

	for remainSteps > 0 {
		// array_shift(s)
		cookie, s = s[0], s[1:]
		chosenIdx = -1
		res = -1
		for idx, child := range g {
			if cookie >= child {
				resTmp = cookie - child
				if resTmp < res || res == -1 {
					res = resTmp
					chosenIdx = idx
				}
			}
		}
		if chosenIdx > -1 {
			if chosenIdx > 0 {
				g = append(g[:chosenIdx], g[chosenIdx+1:]...)
			} else {
				g = g[1:]
			}
			content++
		}
		remainSteps--
	}

	return content
}
