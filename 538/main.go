package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gSum int

func convertBST(root *TreeNode) *TreeNode {
	gSum = 0
	convertBSTInner(root)
	return root
}

func convertBSTInner(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTInner(root.Right)
	root.Val = root.Val + gSum
	gSum = root.Val
	convertBSTInner(root.Left)
}

func convertBST2(root *TreeNode) *TreeNode {
	convertBSTInner2(root, 0)
	return root
}

// method2
func convertBSTInner2(root *TreeNode, rightSum int) int {
	if root == nil {
		return rightSum
	}
	right := convertBSTInner2(root.Right, rightSum)
	root.Val = root.Val + right
	convertBSTInner2(root.Left, root.Val)
	return nearestNodeValue(root)
}

func nearestNodeValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return root.Val
	}
	return nearestNodeValue(root.Left)
}

func inorderedTravel(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	inorderedTravel(root.Left)
	inorderedTravel(root.Right)
}

func main() {
	root := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: 3}

	node3 := &TreeNode{Val: -4}
	node4 := &TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2

	node1.Left = node3
	node1.Right = node4

	//convertBST2(root)
	convertBST(root)
	inorderedTravel(root)

}
