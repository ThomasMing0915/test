package main

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	//stop recursive condition
	if root == nil {
		return 0
	}

	//stop recursive condition
	if root.Left == nil && root.Right == nil {
		return 1
	}

	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)

	return max(depthLeft, depthRight) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
