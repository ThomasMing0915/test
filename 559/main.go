package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

var max int

func maxDepth(root *Node) int {
	max = 0
	maxDepthInner(root)
	return max
}

func maxDepthInner(node *Node) int {
	if node == nil {
		return 0
	}
	subDepth := 0
	for i := 0; i < len(node.Children); i++ {
		depth := maxDepth(node.Children[i])
		if depth > subDepth {
			subDepth = depth
		}
	}
	if subDepth+1 > max {
		max = subDepth + 1
	}
	return subDepth + 1
}

func main() {
	node := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node.Children = append(node.Children, node3)
	node.Children = append(node.Children, node2)
	node.Children = append(node.Children, node4)
	node3.Children = append(node3.Children, node5)
	node3.Children = append(node3.Children, node6)

	fmt.Println(maxDepth(node))
}
