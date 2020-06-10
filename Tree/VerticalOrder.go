package main

import "fmt"

func VeriticalOrder(root *TreeNode, dia *map[int][]int, hd int) {
	if root == nil {
		return
	}

	(*dia)[hd] = append((*dia)[hd], root.data)

	VeriticalOrder(root.left, dia, hd-1)
	VeriticalOrder(root.right, dia, hd+1)

}

func verticalOrder(tree *TreeNode) {
	dia := make(map[int][]int)

	VeriticalOrder(tree, &dia, 0)

	for _, h := range dia {
		fmt.Println(h)
	}
}
