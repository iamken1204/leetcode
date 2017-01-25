package main

import "fmt"
import "github.com/iamken1204/leetcode/problems"

func main() {
	q := []int{1, 2, 3, 2, 3}
	fmt.Println("Q:", q)
	fmt.Println("Ans:", problems.SingleNumber(q))
}
