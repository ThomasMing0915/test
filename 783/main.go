package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var min int
//
//func minDiffInBST(root *TreeNode) int {
//	min = 1000000000000000
//	fackRoot := &TreeNode{Val: 0}
//	fackRoot.Left = root
//	minDiffInBSTInner(fackRoot, "left")
//	return min
//}
//
//func minDiffInBSTInner(root *TreeNode, direction string) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	if root.Left != nil && abs(root.Val, root.Left.Val) < min {
//		min = abs(root.Val, root.Left.Val)
//	}
//
//	if root.Right != nil && abs(root.Val, root.Right.Val) < min {
//		min = abs(root.Val, root.Right.Val)
//	}
//
//	lNode := minDiffInBSTInner(root.Left, "left")
//	rNode := minDiffInBSTInner(root.Right, "right")
//
//	if lNode != nil && abs(root.Val, lNode.Val) < min {
//		min = abs(root.Val, lNode.Val)
//	}
//	if rNode != nil && abs(root.Val, rNode.Val) < min {
//		min = abs(root.Val, rNode.Val)
//	}
//
//	if direction == "left" && rNode != nil {
//		return rNode
//	}
//
//	if direction == "right" && lNode != nil {
//		return lNode
//	}
//	return root
//}

var lastVisitNode *TreeNode
var min int

func minDiffInBST(root *TreeNode) int {
	min = 0x7fffffff
	lastVisitNode = nil

	dfs(root)

	return min
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}

	dfs(node.Left)
	if lastVisitNode != nil {
		diff := abs(lastVisitNode.Val, node.Val)
		if diff < min {
			min = diff
		}
	}
	lastVisitNode = node
	dfs(node.Right)
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	root := &TreeNode{Val: 41}
	node1 := &TreeNode{Val: 19}
	node2 := &TreeNode{Val: 62}
	node3 := &TreeNode{Val: 11}
	node4 := &TreeNode{Val: 31}
	node5 := &TreeNode{Val: 89}
	node6 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 18}
	node8 := &TreeNode{Val: 74}
	node9 := &TreeNode{Val: 20}

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4

	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node7.Left = node9
	node5.Left = node8

	fmt.Println(minDiffInBST(root))
}
