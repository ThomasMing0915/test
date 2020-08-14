package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, b := isBalancedInner(root)
	return b
}

func isBalancedInner(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftDepth, leftBalanced := isBalancedInner(node.Left)
	rightDepth, rightBalanced := isBalancedInner(node.Right)

	if !leftBalanced || !rightBalanced {
		return -1, false
	}

	if absDiff(leftDepth, rightDepth) > 1 {
		return -1, false
	}

	return maxValue(leftDepth, rightDepth) + 1, true
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func maxValue(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4

	fmt.Println(isBalanced(root))

	root = &TreeNode{Val: 1}
	node1 = &TreeNode{Val: 2}
	node2 = &TreeNode{Val: 2}
	node3 = &TreeNode{Val: 3}
	node4 = &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 4}

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node3.Left = node5
	node3.Right = node6
	fmt.Println(isBalanced(root))
}
