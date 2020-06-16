package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0, 32)
	dfs(root, &res)
	return res
}

func dfs(node *Node, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	for i := 0; i < len(node.Children); i++ {
		dfs(node.Children[i], res)
	}
}

func main() {
	root := &Node{Val: 1}
	node3 := &Node{Val: 3}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}

	root.Children = append(root.Children, node3)
	root.Children = append(root.Children, node2)
	root.Children = append(root.Children, node4)

	node3.Children = append(node3.Children, node5)
	node3.Children = append(node3.Children, node6)

	fmt.Println(preorder(root))
}
