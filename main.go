package main

import "fmt"
import "github.com/iamken1204/leetcode/problems"

func main() {

	nodes := &problems.TreeNode{
		Val: 2,
		Left: &problems.TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &problems.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(problems.InvertTree(nodes))

}
