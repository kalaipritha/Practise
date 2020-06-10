package main

import "math"

func height(root *TreeNode) int16 {
	if root == nil {
		return 0
	}

	lheight := height(root.left)
	rheight := height(root.right)

	if lheight > rheight {
		return lheight + 1
	} else {
		return rheight + 1
	}
}

func Diameter(root *TreeNode) int16 {
	if root == nil {
		return 0
	}

	lheight := height(root.left)
	rheight := height(root.right)
	ldiameter := Diameter(root.left)
	rdiameter := Diameter(root.right)
	return int16(math.Max((float64(lheight) + float64(rheight) + 1), math.Max(float64(ldiameter), float64(rdiameter))))
}

