package main

import "fmt"

func InorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	InorderTraversal(root.left)
	fmt.Print(root.data, " ")
	InorderTraversal(root.right)

}

func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	PreorderTraversal(root.left)
	PreorderTraversal(root.right)

}

func PostorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostorderTraversal(root.left)
	PostorderTraversal(root.right)
	fmt.Print(root.data, " ")
}
