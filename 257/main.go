package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0, 32)
	if root == nil {
		return paths
	}
	rootVal := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		paths = append(paths, rootVal)
	}
	if root.Left != nil {
		dfs(root.Left, rootVal)
	}
	if root.Right != nil {
		dfs(root.Right, rootVal)
	}
	return paths
}

func dfs(node *TreeNode, path string) {
	if node == nil {
		return
	}
	path = path + "->" + strconv.Itoa(node.Val)
	//leaf node
	if node.Left == nil && node.Right == nil {
		paths = append(paths, path)
		return
	}
	dfs(node.Left, path)
	dfs(node.Right, path)
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}

	root.Left = node2
	root.Right = node3
	node2.Right = node5

	fmt.Println(binaryTreePaths(root))
}
