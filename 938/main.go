package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum = 0
	dfs(root, L, R)
	return sum
}

func dfs(node *TreeNode, L int, R int) {
	if node == nil {
		return
	}
	if node.Val >= L && node.Val <= R {
		sum += node.Val
	}

	if node.Val < L {
		dfs(node.Right, L, R)
	} else if node.Val > R {
		dfs(node.Left, L, R)
	} else {
		dfs(node.Left, L, R)
		dfs(node.Right, L, R)
	}
}
