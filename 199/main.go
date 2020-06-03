package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res = make([]int, 0, 8)
	var queue = make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelLength := len(queue)
		res = append(res, queue[0].Val)
		for levelLength > 0 {
			node := queue[0]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			queue = queue[1:]
			levelLength--
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}

	root.Left = node2
	root.Right = node3
	node2.Right = node5
	node3.Right = node4

	fmt.Println(rightSideView(root))
}
