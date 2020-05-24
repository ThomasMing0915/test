package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sumLevel map[int]int //表示每次节点的数字之和
var sumCount map[int]int //表示每层节点的数量

func averageOfLevels(root *TreeNode) []float64 {
	sumLevel = make(map[int]int, 128)
	sumCount = make(map[int]int, 128)
	averageOfLevelsInner(root, 0)
	average := make([]float64, len(sumLevel))
	for i := range sumLevel {
		average[i] = float64(sumLevel[i]) / float64(sumCount[i])
	}
	return average
}

func averageOfLevelsInner(root *TreeNode, height int) {
	if root == nil {
		return
	}
	if _, ok := sumCount[height]; !ok {
		sumCount[height] = 1
	} else {
		sumCount[height] += 1
	}
	if _, ok := sumLevel[height]; !ok {
		sumLevel[height] = root.Val
	} else {
		sumLevel[height] += root.Val
	}

	averageOfLevelsInner(root.Left, height+1)
	averageOfLevelsInner(root.Right, height+1)
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

	average := averageOfLevels(root)
	fmt.Println(average)
}
