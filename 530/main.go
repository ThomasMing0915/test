package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var lastVisitNode *TreeNode
var min int

func getMinimumDifference(root *TreeNode) int {
	lastVisitNode = nil
	min = 0x7fffffff
	dfs(root)
	return min
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	if lastVisitNode != nil {
		diff := abs(lastVisitNode.Val, node.Val)
		if diff < min {
			min = diff
		}
	}
	lastVisitNode = node

	dfs(node.Right)
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 9}

	root.Right = node3
	node3.Left = node2

	fmt.Println(getMinimumDifference(root))
}
