package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	//root is nil define return false
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		return true
	}

	if hasPathSum(root.Left, sum-root.Val) {
		return true
	}

	if hasPathSum(root.Right, sum-root.Val) {
		return true
	}

	return false
}
