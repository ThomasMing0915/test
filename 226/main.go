package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invertTreeInner(root)
	return root
}

func invertTreeInner(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeInner(root.Left)
	invertTreeInner(root.Right)
}

func inorderTravel(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	inorderTravel(root.Left)
	inorderTravel(root.Right)
}

func main() {
	root := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node7 := &TreeNode{Val: 7}
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node9 := &TreeNode{Val: 9}

	root.Left = node2
	root.Right = node7
	node2.Left = node1
	node2.Right = node3
	node7.Left = node6
	node7.Right = node9

	invertTree(root)
	inorderTravel(root)
}
