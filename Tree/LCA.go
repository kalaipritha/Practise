package main

import "fmt"

func LowestCommonAncestor(root *TreeNode, tar1 int, tar2 int) {
	var stack1, stack2 Stack
	cur := root
	var LCA1, LCA2 *TreeNode
	stack1_flag := false
	stack2_flag := false

	for cur != nil || len(stack1) > 0 || len(stack2) > 0 {
		for cur != nil {
			if !stack1_flag {
				stack1.Push(cur)
				if cur.data == tar1 {
					stack1_flag = true
				}
			}

			if !stack2_flag {
				stack2.Push(cur)
				if cur.data == tar2 {
					stack2_flag = true
				}
			}

			if !stack1_flag || !stack2_flag {
				cur = cur.left
			} else {
				break
			}
		}
		if !stack1_flag {
			cur = stack1.Pop()
		}

		if !stack2_flag {
			cur = stack2.Pop()
		}

		if stack1_flag && stack2_flag {
			if len(stack1) > len(stack2) {
				for len(stack1) != len(stack2) {
					stack1.Pop()
				}
			}
			if len(stack2) > len(stack1) {
				for len(stack2) != len(stack1) {
					stack2.Pop()
				}
			}
			LCA1 = stack1.Pop()
			LCA2 = stack2.Pop()
			for LCA1.data != LCA2.data && len(stack2) > 0 && len(stack1) > 0 {
				LCA1 = stack1.Pop()
				LCA2 = stack2.Pop()
			}
			break
		}
		if !stack1_flag && stack1.Top().right.data == cur.data {
			cur = nil
		} else if !stack2_flag && stack2.Top().right.data == cur.data {
			cur = nil
		} else {
			if !stack1_flag {
				cur = stack1.Top().right
			} else {
				cur = stack2.Top().right
			}
		}
	}

	if LCA1.data == LCA2.data {
		fmt.Println("The Lowest Common Ancestor is", LCA1.data)
	}
}
