package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	levelMaxValue := make([]int, 0, 32)
	queue := make([]*TreeNode, 0, 8)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelNodeNums := len(queue)
		maxValue := queue[0].Val
		for levelNodeNums > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Val > maxValue {
				maxValue = node.Val
			}
			queue = queue[1:]
			levelNodeNums--
		}
		levelMaxValue = append(levelMaxValue, maxValue)
	}
	return levelMaxValue
}
