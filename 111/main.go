package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return minDepth(root.Right) + 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
