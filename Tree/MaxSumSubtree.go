package main

import "math"

func MaxSumSubtree(root *TreeNode, max_sum *int) int {
	if root == nil {
		return 0
	}

	lt := SumOfSubtree(root.left)
	rt := SumOfSubtree(root.right)
	sum := lt + rt + root.data

	*max_sum = int(math.Max(float64(*max_sum), float64(sum)))

	MaxSumSubtree(root.left, max_sum)
	MaxSumSubtree(root.right, max_sum)
	return *max_sum

}
