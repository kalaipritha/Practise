package main

import "fmt"

func IsSumtree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lt := SumOfSubtree(root.left)
	rt := SumOfSubtree(root.right)

	fmt.Println("the left subtree for ", root.data, " is ", lt)
	fmt.Println("the right subtree for ", root.data, " is ", rt)

	if (root.data == lt+rt) && IsSumtree(root.left) && IsSumtree(root.right) {
		return true
	} else {
		return false
	}

}
