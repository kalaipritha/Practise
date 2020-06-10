package main

import "fmt"

func Display(nodes []int, length int, flag int) {
	var i int

	if flag == 0 {
		for i = length; i >= 0; i-- {
			fmt.Println(nodes[i])
		}
	} else if flag == 1 {
		for i = 0; i <= length; i++ {
			fmt.Println(nodes[i])
		}
	}
}

func PrintPath(root *TreeNode, ht int, flag int, pathlen int, nodes []int) int {
	if root == nil {
		return flag
	}

	nodes[pathlen] = root.data
	if root.left == nil && root.right == nil {

		if pathlen == ht-1 && (flag == 0 || flag == 1) {
			Display(nodes, pathlen, flag)
			return 2
		}
	} else {
		flag = PrintPath(root.left, ht, flag, pathlen+1, nodes)
		flag = PrintPath(root.right, ht, flag, pathlen+1, nodes)

	}
	return flag
}

func CheckPath(root *TreeNode, dia int16) {
	if root == nil {
		return
	}
	lheight := height(root.left)
	rheight := height(root.right)

	if int16(lheight+rheight+1) >= dia {
		a := make([]int, 10)
		PrintPath(root.left, int(lheight), 0, 0, a)
		fmt.Println(root.data)
		PrintPath(root.right, int(rheight), 1, 0, a)
	} else {
		CheckPath(root.left, dia)
		CheckPath(root.right, dia)
	}

}

func Leaf_to_Leaf(root *TreeNode) {
	if root == nil {
		return
	}
	diameter := Diameter(root)

	fmt.Println("the diameter is", diameter)

	CheckPath(root, diameter)
	fmt.Println("______________________________")

}
