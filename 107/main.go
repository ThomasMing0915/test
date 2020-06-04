package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0, 32)
	queue = append(queue, root)
	res := make([][]int, 0, 8)
	for len(queue) > 0 {
		levelLength := len(queue)
		levelSlice := make([]int, 0, levelLength)
		for levelLength > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelSlice = append(levelSlice, node.Val)
			queue = queue[1:]
			levelLength--
		}
		res = append(res, levelSlice)
	}
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func main() {
	root := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}

	root.Left = node9
	root.Right = node20
	node20.Left = node15
	node20.Right = node7

	fmt.Println(levelOrderBottom(root))
}
