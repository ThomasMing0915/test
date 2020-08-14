package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	dfs(root)
	return max
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	if left+right > max {
		max = left + right
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node5.Right = node6

	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5

	fmt.Println(diameterOfBinaryTree(root))
}
