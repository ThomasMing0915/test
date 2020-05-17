package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//binary tree is Binary Search Tree?
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if comRes := method1(root.Left, root.Val, "<") && method1(root.Right, root.Val, ">"); !comRes {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

//recursive from up to down
func method1(node *TreeNode, val int, opr string) bool {
	if node == nil {
		return true
	}

	compRes := false
	switch opr {
	case "<":
		compRes = node.Val < val
	case ">":
		compRes = node.Val > val
	}
	if !compRes {
		return false
	}

	return method1(node.Left, val, opr) && method1(node.Right, val, opr)
}

//error solution why?
func isValidBST2(root *TreeNode) bool {
	return method2(root)
}

//recursive from down to up
func method2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	//check down
	if root.Left != nil && !method2(root.Left) {
		return false
	}
	if root.Right != nil && !method2(root.Right) {
		return false
	}

	//check up
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return true
}

func main() {
	root1 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	root1.Left = node1
	root1.Right = node3

	root2 := &TreeNode{Val: 5}
	node1 = &TreeNode{Val: 1}
	node3 = &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}
	root2.Left = node1
	root2.Right = node4
	node4.Left = node3
	node4.Right = node6

	root3 := &TreeNode{Val: 10}
	node5 := &TreeNode{Val: 5}
	node15 := &TreeNode{Val: 15}
	node6 = &TreeNode{Val: 6}
	node20 := &TreeNode{Val: 20}
	root3.Left = node5
	root3.Right = node15
	node15.Left = node6
	node15.Right = node20

	fmt.Println("method1")
	fmt.Println("root1 is BST", isValidBST(root1))
	fmt.Println("root2 is BST", isValidBST(root2))
	fmt.Println("root3 is BST", isValidBST(root3))

	//fmt.Println("method2")
	//fmt.Println("root1 is BST", isValidBST2(root1))
	//fmt.Println("root2 is BST", isValidBST2(root2))
}
