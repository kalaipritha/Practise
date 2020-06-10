package main

import "fmt"

func TopView(root *TreeNode, dia *map[int]int, hd int) {
	if root == nil {
		return
	}

	_, OK := (*dia)[hd]

	if !OK {
		(*dia)[hd] = root.data
	}

	TopView(root.left, dia, hd-1)
	TopView(root.right, dia, hd+1)

}

func topView(root *TreeNode) {
	dia := make(map[int]int)
	TopView(root, &dia, 0)

	for _, h := range dia {
		fmt.Println(h)
	}
}
