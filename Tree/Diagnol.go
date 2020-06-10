package main

import "fmt"

func Diagnol(root *TreeNode, dia *map[int][]int, hd int) {
	if root == nil {
		return
	}

	(*dia)[hd] = append((*dia)[hd], root.data)

	Diagnol(root.left, dia, hd+1)
	Diagnol(root.right, dia, hd)

}

func SumOfDiagnolElements(root *TreeNode) {
	dia := make(map[int][]int)
	Diagnol(root, &dia, 0)

	for _, h := range dia {
		var sum int
		for _, l := range h {
			sum = sum + l
		}
		fmt.Println(sum)
	}
}
