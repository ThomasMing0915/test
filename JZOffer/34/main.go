package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gres [][]int
var gsum int

func pathSum(root *TreeNode, sum int) [][]int {
	gres = make([][]int, 0, 32)
	path := make([]int, 0, 500)
	gsum = 0

	dfs(root, path, sum)

	return gres
}

func dfs(node *TreeNode, path []int, sum int) {
	if node == nil {
		return
	}
	gsum += node.Val
	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil && gsum == sum {
		tmpPath := make([]int, len(path))
		copy(tmpPath, path)
		gres = append(gres, tmpPath)
	}

	dfs(node.Left, path, sum)
	dfs(node.Right, path, sum)

	//回溯
	gsum -= node.Val
	// path = path[:len(path)-1] 这里为啥不需要回溯，因为是传递的是切换，当函数返回到上一层调用时自动还原了
}

func main() {
	root := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node8 := &TreeNode{Val: 8}
	node11 := &TreeNode{Val: 11}
	node13 := &TreeNode{Val: 13}
	node42 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node2 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}

	root.Left = node4
	root.Right = node8
	node4.Left = node11
	node8.Left = node13
	node8.Right = node42
	node11.Left = node7
	node11.Right = node2
	node42.Left = node5
	node42.Right = node1

	res := pathSum(root, 22)
	fmt.Println(res)
}
