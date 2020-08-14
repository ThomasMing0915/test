package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0, 1)
	res := make([][]int, 0, 32)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelCount := len(queue)
		levelNodeValues := make([]int, 0, levelCount)
		for levelCount > 0 {
			node := queue[0]
			for i := 0; i < len(node.Children); i++ {
				queue = append(queue, node.Children[i])
			}

			levelNodeValues = append(levelNodeValues, node.Val)
			queue = queue[1:]
			levelCount--
		}
		res = append(res, levelNodeValues)
	}
	return res
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

	fmt.Println(levelOrder(root))
}
