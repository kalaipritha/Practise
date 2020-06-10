package main

import (
	"fmt"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

type Stack []*TreeNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str *TreeNode) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Top() *TreeNode {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	return element
}

func (s *Stack) Pop() *TreeNode {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

func initTree() *Tree {
	return &Tree{
		root: &TreeNode{},
	}
}

func (t *TreeNode) insert(data int) {
	if t.left == nil {
		temp := &TreeNode{}
		temp.data = data
		temp.left = nil
		temp.right = nil
		t.left = temp
	} else {
		temp := &TreeNode{}
		temp.data = data
		temp.left = nil
		temp.right = nil
		t.right = temp
	}
}

func main() {
	tree := initTree()
	tree.root.data = 1
	tree.root.insert(2)
	tree.root.insert(3)
	tree.root.left.insert(9)
	tree.root.left.insert(6)
	tree.root.left.left.insert(0)
	tree.root.left.left.insert(10)
	tree.root.left.left.right.insert(23)
	tree.root.left.left.right.left.insert(24)
	tree.root.left.left.right.left.insert(25)

	tree.root.left.right.insert(11)
	tree.root.left.right.left.insert(20)
	tree.root.left.right.left.insert(21)
	tree.root.left.right.left.left.insert(29)
	tree.root.right.insert(4)
	tree.root.right.insert(5)
	tree.root.right.left.insert(12)
	tree.root.right.left.insert(7)

	tree.root.left.right.insert(1)
	tree.root.left.right.insert(12)
	tree.root.left.right.right.insert(0)
	tree.root.left.right.right.insert(2)
	tree.root.right.insert(7)
	InorderTraversal(tree.root)
	PreorderTraversal(tree.root)
	PostorderTraversal(tree.root)
	var sum int
	SumAllNodes(tree.root, &sum)
	fmt.Println("the sum is", sum)
	var stack Stack
	PrintAllPaths(tree.root, stack)
	HasPathSum(tree.root, stack, 13)

	fmt.Println(IsSumtree(tree.root))
	fmt.Println(SumSubTree(tree.root, 23))
	fmt.Println(MaxSumSubtree(tree.root, &sum))
	FindDistance(tree.root, 4, 0, &sum)
	fmt.Println(sum)

	LowestCommonAncestor(tree.root, -2, 3)

	InorderSuccessor(tree.root, 3)
	InorderPredecessor(tree.root, 3)
	InorderSuccessor_Predeccessor(tree.root, 3)

	dia := Diameter(tree.root)
	fmt.Println("the diameter is ", dia)

	verticalOrder(tree.root)
	topView(tree.root)

	SumOfDiagnolElements(tree.root)

	fmt.Println(Level(tree.root, 0, 11))

	Leaf_to_Leaf(tree.root)
}
