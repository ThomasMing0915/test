package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0, 32)
	res := make([][]int, 0, 8)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		level++
		levelLength := len(queue)
		levelSlice := make([]int, 0, 8)
		if level%2 == 1 {
			for i := 0; i < levelLength; i++ {
				levelSlice = append(levelSlice, queue[i].Val)
			}
		} else {
			for i := levelLength - 1; i >= 0; i-- {
				levelSlice = append(levelSlice, queue[i].Val)
			}
		}

		for levelLength > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			levelLength--
		}
		res = append(res, levelSlice)
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
	node2.Left = node4
	node3.Right = node5

	fmt.Println(zigzagLevelOrder(root))
}
