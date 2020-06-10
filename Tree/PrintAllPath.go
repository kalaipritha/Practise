package main

import "fmt"

func PrintAllPaths(root *TreeNode, Stack Stack) {
	Stack.Push(root)

	if root.left == nil && root.right == nil {
		for !Stack.IsEmpty() {
			s := Stack.Pop()
			fmt.Print(s.data, " -> ")
		}
		fmt.Println()
		return
	}

	PrintAllPaths(root.left, Stack)
	PrintAllPaths(root.right, Stack)

}
