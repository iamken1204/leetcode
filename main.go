package main

import "fmt"
import "github.com/iamken1204/leetcode/problems"

func main() {

	bbb := &problems.TreeNode{}
	aaa := &problems.TreeNode{Right: bbb}
	node := &problems.TreeNode{
		Right: aaa,
	}

	fmt.Println(problems.MaxDepth(node))

}
