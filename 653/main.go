package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func findTarget(root *TreeNode, k int) bool {
	res = make([]int, 0, 32)
	dfs(root, &res)

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i]+res[j] == k {
				return true
			}
		}
	}
	return false
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	dfs(node.Left, res)
	dfs(node.Right, res)
}
