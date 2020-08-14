package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var stack *Stack

func flatten(root *TreeNode) {
	stack = New()
	flattenInner(root)
}

func flattenInner(root *TreeNode) {
	//stop recursive condition
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if item := stack.Pop(); item != nil {
			if node, ok := item.(*TreeNode); ok {
				root.Right = node
				flattenInner(node)
			} else {
				panic("item is not treenode")
			}
		} else {
			return
		}
	}

	if root.Left != nil && root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		flattenInner(root.Right)
	}
	if root.Left == nil && root.Right != nil {
		flattenInner(root.Right)
	}
	if root.Left != nil && root.Right != nil {
		stack.Push(root.Right)
		root.Right = root.Left
		root.Left = nil
		flattenInner(root.Right)
	}
}

func InorderTravel(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	//InorderTravel(root.Left)
	if root.Left != nil {
		panic("left is not nil")
	}
	InorderTravel(root.Right)

}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}

	root.Left = node2
	root.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Right = node6

	flatten(root)
	InorderTravel(root)

	root2 := &TreeNode{Val: 1}
	node21 := &TreeNode{Val: 2}
	node22 := &TreeNode{Val: 3}
	root2.Left = node21
	node21.Left = node22
	flatten(root2)
	InorderTravel(root2)

	root3 := &TreeNode{Val: 1}
	node31 := &TreeNode{Val: 2}
	node32 := &TreeNode{Val: 3}
	node33 := &TreeNode{Val: 4}
	root3.Left = node31
	root3.Right = node32
	node32.Left = node33
	flatten(root3)
	InorderTravel(root3)

	flatten(nil)
}
