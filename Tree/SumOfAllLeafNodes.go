package main

func SumAllNodes(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return root.data
	}
	*sum = *sum + SumAllNodes(root.left, sum) + SumAllNodes(root.right, sum)
	return 0
}
