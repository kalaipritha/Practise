package main

import "fmt"

func SumSubTree(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	lt := SumOfSubtree(root.left)
	rt := SumOfSubtree(root.right)

	fmt.Println("the left subtree for ", root.data, " is ", lt)
	fmt.Println("the right subtree for ", root.data, " is ", rt)

	return (target == lt) || (target == rt) || SumSubTree(root.left, target) || SumSubTree(root.right, target)
}
