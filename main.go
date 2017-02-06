package main

import "fmt"
import "github.com/iamken1204/leetcode/problems"

func main() {

	children := []int{1, 2, 3, 5}
	cookies := []int{3, 1, 2}

	fmt.Println(problems.FindContentChildren(children, cookies))

}
