package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesInner(root)
}

func sumOfLeftLeavesInner(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeavesInner(root.Right)
	}

	return sumOfLeftLeavesInner(root.Left) + sumOfLeftLeavesInner(root.Right)
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

	root2 := &TreeNode{Val: 1}
	node21 := &TreeNode{Val: 1}
	node22 := &TreeNode{Val: 2}
	node23 := &TreeNode{Val: 3}
	node24 := &TreeNode{Val: 4}
	node25 := &TreeNode{Val: 5}
	node26 := &TreeNode{Val: 6}
	//root2.Left = node21
	//root2.Right = node22
	//node21.Left = node23
	//node21.Right = node24
	//node22.Left = node25
	//node22.Right = node26

	root2.Right = node21
	node21.Right = node22
	node22.Right = node23
	node23.Right = node24
	node24.Right = node25
	node25.Right = node26

	fmt.Println(sumOfLeftLeaves(root))
	fmt.Println(sumOfLeftLeaves(root2))
}
