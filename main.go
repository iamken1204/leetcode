package main

import "fmt"
import "github.com/iamken1204/leetcode/problems"

func main() {

	children := []int{10, 9, 8, 7}
	cookies := []int{7, 8}

	fmt.Println(problems.FindContentChildren(children, cookies))

}
