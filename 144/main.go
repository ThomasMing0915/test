package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(root *TreeNode) []int {
	arr := make([]int, 0, 32)
	preorderTravel(root, &arr)
	return arr
}

//recursive algorithm
func preorderTravel(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	*arr = append(*arr, node.Val)
	preorderTravel(node.Left, arr)
	preorderTravel(node.Right, arr)
}

// iteration algorithm : use stack structure
func preorderTraversal2(root *TreeNode) []int {
	arr := make([]int, 0, 32)
	stack := New()
	if root == nil {
		return arr
	}
	stack.Push(root)

	for !stack.IsEmpty() {
		item := stack.Pop()
		popTreeNode, ok := item.(*TreeNode)
		if !ok {
			panic("pop item is not tree node")
		}
		arr = append(arr, popTreeNode.Val)
		if popTreeNode.Right != nil {
			stack.Push(popTreeNode.Right)
		}
		if popTreeNode.Left != nil {
			stack.Push(popTreeNode.Left)
		}
	}

	return arr
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	root.Right = node2
	node2.Left = node3

	arr := preorderTraversal2(root)
	for _, v := range arr {
		fmt.Println(v)
	}
}
