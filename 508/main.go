package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gSumMap map[int]int //key为节点和，value为出现的次数

func findFrequentTreeSum(root *TreeNode) []int {
	gSumMap = make(map[int]int, 32)
	dfs(root)

	//统计gSumMap中元素出现次数最多的那个是多少
	maxAppearCount := 0
	for _, v := range gSumMap {
		if v > maxAppearCount {
			maxAppearCount = v
		}
	}

	res := make([]int, 0, len(gSumMap))
	//挑出出现次数最多的那些元素
	for k, v := range gSumMap {
		if v == maxAppearCount {
			res = append(res, k)
		}
	}
	return res
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSum := dfs(node.Left)
	rightSum := dfs(node.Right)
	sum := leftSum + rightSum + node.Val

	if _, ok := gSumMap[sum]; ok {
		gSumMap[sum] += 1
	} else {
		gSumMap[sum] = 1
	}

	return sum
}

func main() {
	root := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: -5}

	root.Left = node2
	root.Right = node3

	fmt.Println(findFrequentTreeSum(root))
}
