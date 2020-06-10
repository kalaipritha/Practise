package main

import "fmt"

func InorderSuccessor(root *TreeNode, tar int) {
	if root == nil {
		return
	}
	var stack Stack

	flag := false
	var successor int
	for root != nil || stack != nil {
		for root != nil {
			stack.Push(root)
			root = root.left
		}

		root = stack.Pop()
		successor = root.data
		if flag == true {
			fmt.Println("the successor is ", successor)
			return
		}

		if successor == tar {
			flag = true
		}

		root = root.right
	}
}

func InorderPredecessor(root *TreeNode, tar int) {
	if root == nil {
		return
	}
	var stack Stack
	var predecessor int
	for root != nil || stack != nil {
		for root != nil {
			stack.Push(root)
			root = root.left
		}

		root = stack.Pop()

		if root.data == tar {
			fmt.Println("the predecessor is", predecessor)
			return
		}
		predecessor = root.data

		root = root.right
	}
}

func InorderSuccessor_Predeccessor(root *TreeNode, tar int) {
	if root == nil {
		return
	}
	var stack Stack
	var predecessor, middle, successor int
	for root != nil || stack != nil {
		for root != nil {
			stack.Push(root)
			root = root.left
		}

		root = stack.Pop()

		if predecessor == 0 {
			predecessor = root.data
		} else if successor == 0 {
			successor = root.data
		} else if middle == 0 {
			middle = root.data
		} else {
			predecessor = middle
			middle = successor
			successor = root.data
		}

		if middle == tar {
			fmt.Println("the predecessor is", predecessor)
			fmt.Println("the successor is ", successor)
			return
		}
		root = root.right
	}
}
