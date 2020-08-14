package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: 0}
	createBinaryTree(root, nums)

	return root
}

func createBinaryTree(node *TreeNode, nums []int) {
	if len(nums) == 0 {
		return
	}

	index := maxValueIndex(nums)
	node.Val = nums[index]

	if len(nums[0:index]) > 0 {
		node.Left = &TreeNode{Val: 0}
		createBinaryTree(node.Left, nums[0:index])
	}

	if len(nums[index+1:]) > 0 {
		node.Right = &TreeNode{Val: 0}
		createBinaryTree(node.Right, nums[index+1:])
	}

}

func maxValueIndex(nums []int) int {
	maxValue := nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxIndex = i
		}
	}
	return maxIndex
}

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	dfs(root)
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	dfs(node.Left)
	dfs(node.Right)
}
