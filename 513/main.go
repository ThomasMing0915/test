package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxHeight int
var maxValue int

func findBottomLeftValue(root *TreeNode) int {
	maxHeight = -1
	dfs(root, 0)
	return maxValue
}

func dfs(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	dfs(node.Left, depth+1)

	//node is leaf node and depth is max
	if node.Left == nil && node.Right == nil && depth > maxHeight {
		maxHeight = depth
		maxValue = node.Val
	}

	dfs(node.Right, depth+1)
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node2.Left = node4
	node3.Left = node5
	node3.Right = node6
	node5.Left = node7

	root.Left = node2
	root.Right = node3
	fmt.Println(findBottomLeftValue(root))
}
