package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var kth int

func kthSmallest(root *TreeNode, k int) int {
	kth = 0
	var kthValue = new(int)

	dfs(root, k, kthValue)

	return *kthValue
}

func dfs(node *TreeNode, k int, kthValue *int) {
	if node == nil {
		return
	}
	dfs(node.Left, k, kthValue)
	kth++
	if kth == k {
		*kthValue = node.Val
		return
	}
	dfs(node.Right, k, kthValue)
}

func main() {
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}

	root.Left = node1
	root.Right = node4
	node1.Right = node2

	fmt.Println(kthSmallest(root, 4))
}
