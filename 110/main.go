package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//check a binary tree is balanced?
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if diff(depth(root.Left), depth(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := depth(node.Left)
	rightDepth := depth(node.Right)

	return max(leftDepth, rightDepth) + 1
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root1 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	root1.Left = node9
	root1.Right = node20
	node20.Left = node15
	node20.Right = node7

	root2 := &TreeNode{Val: 1}
	node2_1 := &TreeNode{Val: 2}
	node2_2 := &TreeNode{Val: 2}
	node3_1 := &TreeNode{Val: 3}
	node3_2 := &TreeNode{Val: 3}
	node4_1 := &TreeNode{Val: 4}
	node4_2 := &TreeNode{Val: 4}
	root2.Left = node2_1
	root2.Right = node2_2
	node2_1.Left = node3_1
	node2_1.Right = node3_2
	node3_1.Left = node4_1
	node3_2.Right = node4_2

	root3 := &TreeNode{Val: 0}

	root4 := &TreeNode{Val: 1}
	root4.Left = &TreeNode{Val: 2}
	root4.Right = &TreeNode{Val: 3}

	root5 := &TreeNode{Val: 1}
	node50 := &TreeNode{Val: 50}
	node51 := &TreeNode{Val: 51}
	root5.Left = node50
	node50.Left = node51

	root6 := &TreeNode{Val: 6}
	node60 := &TreeNode{Val: 60}
	node61 := &TreeNode{Val: 61}
	root6.Right = node60
	node60.Right = node61

	fmt.Println(isBalanced(root1))
	fmt.Println(isBalanced(root2))
	fmt.Println(isBalanced(root3))
	fmt.Println(isBalanced(root4))
	fmt.Println(isBalanced(root5))
	fmt.Println(isBalanced(root6))
}
