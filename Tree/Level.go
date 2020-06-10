package main

func Level(root *TreeNode, level int, tar int) int {
	if root == nil {
		return 0
	}

	if root.data == tar {
		return level
	}
	var lev int
	lev = Level(root.left, level+1, tar)
	if lev != 0 {
		return lev
	}
	lev = Level(root.right, level+1, tar)
	return lev
}
