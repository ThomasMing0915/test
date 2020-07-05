package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 32)
	queue := make([]*TreeNode, 0, 500)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelNodeNums := len(queue)
		levelNode := make([]int, 0, levelNodeNums)
		for levelNodeNums > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelNode = append(levelNode, node.Val)
			queue = queue[1:]
			levelNodeNums--
		}
		res = append(res, levelNode)
	}
	return res
}
