package main

func FindDistance(root *TreeNode, target int, dis int, cur *int) {
	if root == nil {
		return
	}

	if root.data == target {
		*cur = dis
	}
	FindDistance(root.left, target, dis+1, cur)
	FindDistance(root.right, target, dis+1, cur)
}
