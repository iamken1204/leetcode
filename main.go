package main

import "fmt"
import "github.com/iamken1204/leetcode/problems"

func main() {

	nums := []int{0, 0, 0, 1, 2, 3, 4}
	problems.MoveZeroes(nums)
	fmt.Println(nums)

}
