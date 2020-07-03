package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一 dfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1) //可以调整为max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//方法二 bfs
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := make([]*TreeNode, 0, 32)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelCount := len(queue) //queue里面装的是某层的节点
		for levelCount > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			levelCount--
		}
		depth++
	}
	return depth
}

//方法三 递归中统计节点数 + 回溯
var gMaxDepth int
var gCurDepth int

func maxDepth3(root *TreeNode) int {
	gMaxDepth, gCurDepth = 0, 0
	dfs(root)
	return gMaxDepth
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	gCurDepth++
	if gCurDepth > gMaxDepth {
		gMaxDepth = gCurDepth
	}
	dfs(node.Left)
	dfs(node.Right)
	gCurDepth--

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

	fmt.Println("方法1:", maxDepth(root))
	fmt.Println("方法2:", maxDepth2(root))
	fmt.Println("方法3:", maxDepth3(root))
}
