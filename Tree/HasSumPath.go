package main

import "fmt"

func HasPathSum(root *TreeNode, Stack Stack, target int) {
	Stack.Push(root)

	if root.left == nil && root.right == nil {
		var sum int
		for !Stack.IsEmpty() {
			s := Stack.Pop()
			sum = sum + s.data
		}
		if sum == target {
			fmt.Println("There is a root to leaf path with sum ", target)
		}
		return
	}

	HasPathSum(root.left, Stack, target)
	HasPathSum(root.right, Stack, target)
}
