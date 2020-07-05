package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ld := dfs(root.Left)
	rd := dfs(root.Right)
	diff := ld - rd
	if diff < -1 || diff > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(dfs(node.Left), dfs(node.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//方法2 自底向上
func isBalanced2(root *TreeNode) bool {
	_, b := recursiveBottom2Up(root)
	return b
}

func recursiveBottom2Up(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	dl, bl := recursiveBottom2Up(node.Left)
	if !bl {
		return 0, bl
	}
	dr, br := recursiveBottom2Up(node.Right)
	if !br {
		return 0, br
	}

	diff := dl - dr
	if diff < -1 || diff > 1 {
		return 0, false
	}
	return max(dl, dr) + 1, true
}

func main() {
	root := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}

	root.Left = node9
	root.Right = node20
	node20.Left = node15
	node20.Right = node7

	fmt.Println(isBalanced(root))
	fmt.Println(isBalanced2(root))

	root2 := &TreeNode{Val: 1}
	node21 := &TreeNode{Val: 2}
	node22 := &TreeNode{Val: 2}
	node31 := &TreeNode{Val: 3}
	node32 := &TreeNode{Val: 3}
	node41 := &TreeNode{Val: 4}
	node42 := &TreeNode{Val: 4}

	root2.Left = node21
	root2.Right = node22
	node21.Left = node31
	node21.Right = node32
	node31.Left = node41
	node31.Right = node42

	fmt.Println(isBalanced(root2))
	fmt.Println(isBalanced2(root2))
}
