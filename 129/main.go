package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	sum       int
	prefixSum int
)

func sumNumbers(root *TreeNode) int {
	sum = 0
	prefixSum = 0

	sumNumbersInner(root)
	return sum
}

func sumNumbersInner(root *TreeNode) {
	if root == nil {
		return
	}
	prefixSum = prefixSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		sum += prefixSum
		prefixSum = prefixSum / 10
		return
	}

	sumNumbersInner(root.Left)
	sumNumbersInner(root.Right)
	prefixSum = prefixSum / 10
}

func main() {
	root := &TreeNode{Val: 1}
	//node1 := &TreeNode{Val: 2}
	//node2 := &TreeNode{Val: 3}
	//root.Right = node1
	//node1.Left = node2

	fmt.Println(sumNumbers(root))
}
