package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if pruneTreeInner(root) {
		return nil
	}
	return root
}

func pruneTreeInner(node *TreeNode) bool {
	if node == nil {
		return true
	}

	lb := pruneTreeInner(node.Left)
	rb := pruneTreeInner(node.Right)
	if lb {
		node.Left = nil
	}
	if rb {
		node.Right = nil
	}
	if node.Val == 0 && lb && rb {
		return true
	}
	return false
}

func midOrderTravel(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	midOrderTravel(root.Left)
	midOrderTravel(root.Right)
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 0}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 0}
	node6 := &TreeNode{Val: 1}
	node7 := &TreeNode{Val: 0}

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	node3.Left = node7

	r := pruneTree(root)
	midOrderTravel(r)
}
