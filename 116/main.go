package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	connectInner(root)
	return root
}

func connectInner(root *Node) {
	if root == nil {
		return
	}
	//leaf node
	if root.Left == nil && root.Right == nil {
		return
	}

	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connectInner(root.Left)
	connectInner(root.Right)
}

func travelNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("\ncurNode val :%d ", node.Val)
	if node.Left != nil {
		fmt.Printf(" leftNode val :%d ", node.Left.Val)
	}
	if node.Right != nil {
		fmt.Printf(" rightNode val :%d ", node.Right.Val)
	}
	if node.Next != nil {
		fmt.Printf(" nextNode val :%d ", node.Next.Val)
	}
	travelNode(node.Left)
	travelNode(node.Right)
}

func main() {
	root := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}

	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	connect(root)
	travelNode(root)
}
