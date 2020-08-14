package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil || (root.Val == key && root.Left == nil && root.Right == nil) {
		return nil
	}

	if root.Val == key && root.Left != nil {
		newKey, _ := adjust(root.Left, "left")
		root.Val = newKey
		return root
	}

	if root.Val == key && root.Right != nil {
		newKey, _ := adjust(root.Right, "right")
		root.Val = newKey
		return root
	}

	midOrderTravel(root.Left, root, key, "left")
	midOrderTravel(root.Right, root, key, "right")
	return root
}

func midOrderTravel(node *TreeNode, nodeParent *TreeNode, key int, direction string) bool {
	if node == nil {
		return false
	}
	if node.Val == key {
		if node.Left != nil {
			newKey, b := adjust(node.Left, "left")
			node.Val = newKey
			if !b {
				node.Left = node.Left.Left
			}

		} else if node.Right != nil {
			newKey, b := adjust(node.Right, "right")
			node.Val = newKey
			if !b {
				node.Right = node.Right.Right
			}

		} else { //要删除的节点是叶子节点
			if direction == "left" {
				nodeParent.Left = nil
			} else if direction == "right" {
				nodeParent.Right = nil
			}

		}
		return true
	}
	if midOrderTravel(node.Left, node, key, "left") {
		return true
	}
	if midOrderTravel(node.Right, node, key, "right") {
		return true
	}
	return false
}

func adjust(node *TreeNode, direction string) (int, bool) {
	if node == nil {
		panic("node is nil")
	}
	p, q := node, node
	for p != nil {
		if direction == "left" {
			if p.Right != nil {
				q = p
				p = p.Right
			} else {
				break
			}
		} else if direction == "right" {
			if p.Left != nil {
				q = p
				p = p.Left
			} else {
				break
			}
		}
	}
	if p != q {
		if direction == "left" {
			q.Right = p.Left
		} else if direction == "right" {
			q.Left = p.Right
		}
		return p.Val, true
	}
	return p.Val, false
}

func preTravel(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	preTravel(node.Left)
	preTravel(node.Right)
}

func main() {
	root := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}

	root.Left = node3
	root.Right = node6
	node3.Left = node2
	node3.Right = node4

	node6.Right = node7

	node9 := &TreeNode{Val: 9}
	node10 := &TreeNode{Val: 10}
	node11 := &TreeNode{Val: 11}
	node6.Left = node9
	node9.Left = node10
	node10.Right = node11

	newRoot := deleteNode(root, 5)
	preTravel(newRoot)
}
