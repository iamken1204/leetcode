package problems

import "math"

// ConstructRectangle represents
// Easy 492. Construct the Rectangle
// https://leetcode.com/problems/construct-the-rectangle/
func ConstructRectangle(area int) []int {
	W := int(math.Sqrt(float64(area)))
	for area%W != 0 {
		W--
	}
	return []int{area / W, W}
}
