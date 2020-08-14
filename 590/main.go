package main

type Node struct {
	Val      int
	Children []*Node
}

var gSlice []int

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	gSlice = make([]int, 0, 32)
	dfs(root)
	return gSlice
}

func dfs(node *Node) {
	if node == nil {
		return
	}
	for i := range node.Children {
		dfs(node.Children[i])
	}
	gSlice = append(gSlice, node.Val)
}
