package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gK int
var gV int

func kthLargest(root *TreeNode, k int) int {
	gK = 0
	gV = 0

	dfs(root, k)
	return gV
}

func dfs(node *TreeNode, k int) {
	if node == nil {
		return
	}
	dfs(node.Right, k)
	gK++
	if gK == k {
		gV = node.Val
		return
	}
	dfs(node.Left, k)
}
