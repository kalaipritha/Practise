package main

func SumOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return (SumOfSubtree(root.left) + root.data + SumOfSubtree(root.right))
}
