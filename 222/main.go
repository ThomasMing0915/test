package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var count int

//func countNodes(root *TreeNode) int {
//	count = 0
//	countNodesInner(root)
//	return count
//}
//
//func countNodesInner(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	count += 1
//	countNodesInner(root.Left)
//	countNodesInner(root.Right)
//}

func leftHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return leftHeight(root.Left) + 1
}

func rightHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return rightHeight(root.Right) + 1
}

func countLevelNodes(root *TreeNode, currentLevel, targetLevel int) int {
	if root == nil {
		return 0
	}
	currentLevel += 1
	if currentLevel == targetLevel {
		return countLeafNodes(root)
	}
	if currentLevel < targetLevel {
		return countLevelNodes(root.Left, currentLevel, targetLevel) + countLevelNodes(root.Right, currentLevel, targetLevel)
	}
	//不会走的这里
	return 0
}

func countNodes2(root *TreeNode) int {
	lh := leftHeight(root.Left) + 1
	rh := rightHeight(root.Right) + 1
	if lh == rh { //满二叉树 直接计算
		return int(math.Pow(2, float64(lh)) - 1)
	}
	if lh < rh {
		panic("left height bigger than right height")
	}

	return countLevelNodes(root, 0, lh-1) + int(math.Pow(2, float64(lh-1))-1)
}

func countLeafNodes(node *TreeNode) int {
	count := 0
	if node.Left != nil {
		count += 1
	}
	if node.Right != nil {
		count += 1
	}
	return count
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6

	root1 := &TreeNode{Val: 1}

	node12 := &TreeNode{Val: 2}
	node13 := &TreeNode{Val: 3}
	node21 := &TreeNode{Val: 4}
	root1.Left = node12
	root1.Right = node13
	node12.Left = node21
	fmt.Println(countNodes2(root))
	fmt.Println(countNodes2(root1))
}
