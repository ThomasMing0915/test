package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var countMap map[string]int
var resMap map[*TreeNode]struct{}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	countMap = make(map[string]int)
	resMap = make(map[*TreeNode]struct{})

	dfs(root, countMap, resMap)

	res := make([]*TreeNode, 0, len(resMap))
	for k := range resMap {
		res = append(res, k)
	}

	return res
}

func dfs(node *TreeNode, countMap map[string]int, resMap map[*TreeNode]struct{}) string {
	if node == nil {
		return "null"
	}

	str := strconv.Itoa(node.Val) + "," + dfs(node.Left, countMap, resMap) + "," + dfs(node.Right, countMap, resMap)
	if count, ok := countMap[str]; ok && count == 1 {
		//fmt.Println(str)
		resMap[node] = struct{}{}
	}
	countMap[str] += 1
	//fmt.Println("---", str)

	return str
}

func main() {
	root := &TreeNode{Val: 2}
	node21 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 11}
	node41 := &TreeNode{Val: 11}
	//node42 := &TreeNode{Val: 4}
	node22 := &TreeNode{Val: 1}
	//node43 := &TreeNode{Val: 0}

	root.Left = node21
	root.Right = node3
	node21.Left = node41
	node3.Left = node22
	//node3.Right = node42
	//node22.Right = node43

	res := findDuplicateSubtrees(root)
	for _, node := range res {
		fmt.Println(node.Val)
	}
}
