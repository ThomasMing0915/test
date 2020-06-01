package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	if root.Val == x || root.Val == y {
		return false
	}
	xval, xdepth := depthKNodeParent(root, x, 0)
	yval, ydepth := depthKNodeParent(root, y, 0)
	if xdepth == ydepth && xval != yval {
		return true
	}
	return false
}

func depthKNodeParent(node *TreeNode, val int, depth int) (int, int) {
	if node == nil {
		return -1, -1
	}
	if node.Left != nil && node.Left.Val == val {
		return node.Val, depth + 1
	}
	if node.Right != nil && node.Right.Val == val {
		return node.Val, depth + 1
	}

	lval, ldepth := depthKNodeParent(node.Left, val, depth+1)
	if lval != -1 && ldepth != -1 {
		return lval, ldepth
	}
	rval, rdepth := depthKNodeParent(node.Right, val, depth+1)
	if rval != -1 && rdepth != -1 {
		return rval, rdepth
	}
	return -1, -1
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	//node5 := &TreeNode{Val: 5}
	root.Left = node2
	root.Right = node3
	node2.Right = node4
	//node3.Right = node5

	fmt.Println(isCousins(root, 2, 3))
}
