package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pathp, pathq, path []*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathp = make([]*TreeNode, 0)
	pathq = make([]*TreeNode, 0)
	path = make([]*TreeNode, 0, 32)

	dfs(root, p, q)

	var i int
	for i = 0; i < len(pathp) && i < len(pathq); i++ {
		if pathp[i] != pathq[i] {
			break
		}
	}
	if i < len(pathp) && i < len(pathq) {
		return pathp[i-1]
	}
	if i >= len(pathp) && len(pathp) > 0 {
		return pathp[len(pathp)-1]
	}
	if i >= len(pathq) && len(pathq) > 0 {
		return pathq[len(pathq)-1]
	}
	return nil
}

func dfs(node *TreeNode, p, q *TreeNode) {
	if node == nil {
		return
	}
	path = append(path, node)
	if p == node {
		pathp = make([]*TreeNode, len(path))
		copy(pathp, path)
	}
	if q == node {
		pathq = make([]*TreeNode, len(path))
		copy(pathq, path)
	}
	dfs(node.Left, p, q)
	dfs(node.Right, p, q)
	path = path[:len(path)-1]
}

func main() {
	root := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}

	root.Left = node5
	root.Right = node1
	node5.Left = node6
	node5.Right = node2
	node1.Left = node0
	node1.Right = node8
	node2.Left = node7
	node2.Right = node4

	p := lowestCommonAncestor(root, node0, node1)
	if p != nil {
		fmt.Println(p.Val)
	}
}
