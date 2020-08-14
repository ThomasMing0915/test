package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return isUnivalTreeInner(root, root.Val)
}

func isUnivalTreeInner(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return isUnivalTreeInner(node.Left, val) && isUnivalTreeInner(node.Right, val)
}
