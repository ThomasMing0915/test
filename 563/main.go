package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func findTilt(root *TreeNode) int {
	sum = 0
	findTiltInner(root)
	return sum
}

func findTiltInner(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lsum := findTiltInner(node.Left)
	rsum := findTiltInner(node.Right)
	sum += absDiff(lsum, rsum)
	return lsum + rsum + node.Val
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
