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

	queue := make([]*TreeNode, 0, 500)
	res := make([][]int, 0, 32)

	queue = append(queue, root)
	flag := 0
	for len(queue) > 0 {
		flag++
		levelNodeCount := len(queue)
		levelNodeValue := make([]int, 0, levelNodeCount)
		for levelNodeCount > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			levelNodeValue = append(levelNodeValue, node.Val)
			queue = queue[1:]
			levelNodeCount--
		}
		if flag%2 == 1 {
			res = append(res, levelNodeValue)
		} else {
			newLevelNodeCount := len(levelNodeValue)
			reverseLevelNodeValue := make([]int, 0, newLevelNodeCount)
			for i := range levelNodeValue {
				reverseLevelNodeValue = append(reverseLevelNodeValue, levelNodeValue[newLevelNodeCount-i-1])
			}
			res = append(res, reverseLevelNodeValue)
		}
	}

	return res
}
