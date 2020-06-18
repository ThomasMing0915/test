package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gv int

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		newRoot := &TreeNode{Val: v}
		newRoot.Left = root
		return newRoot
	}

	gv = v
	dfs(root, 1, d-1)
	return root
}

func dfs(node *TreeNode, depth int, d int) {
	if node == nil {
		return
	}
	if depth == d {
		tmpLeftNode := node.Left
		tmpRightNode := node.Right

		node.Left = &TreeNode{Val: gv}
		node.Left.Left = tmpLeftNode

		node.Right = &TreeNode{Val: gv}
		node.Right.Right = tmpRightNode

		return
	}
	dfs(node.Left, depth+1, d)
	dfs(node.Right, depth+1, d)
}

func midOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	midOrder(node.Left)
	midOrder(node.Right)
}

func main() {
	root := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node6 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 5}

	root.Left = node2
	root.Right = node6
	node2.Left = node3
	node2.Right = node1
	node6.Left = node5

	newRoot := addOneRow(root, 1, 2)
	midOrder(newRoot)
}
