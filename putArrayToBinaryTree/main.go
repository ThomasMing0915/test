package main

import "fmt"

//question how to put an ordered array to binary tree?

func CreateBTreeByOrderedArray(orderedArray []int) *TreeNode {
	if len(orderedArray) == 0 {
		return nil
	}

	root := &TreeNode{}
	createBTreeByOrderedArray(root, orderedArray, 0, len(orderedArray)-1)

	return root
}

func createBTreeByOrderedArray(treeNode *TreeNode, array []int, left, right int) {
	mid := left + (right-left)/2
	treeNode.Val = array[mid]

	//左子树的数组的起始终止索引
	leftPartLeft := left
	leftPartRight := mid - 1

	rightPartLeft := mid + 1
	rightPartRight := right

	if leftPartLeft <= leftPartRight {
		leftPartRoot := &TreeNode{}
		treeNode.Left = leftPartRoot
		createBTreeByOrderedArray(leftPartRoot, array, leftPartLeft, leftPartRight)
	}

	if rightPartLeft <= rightPartRight {
		rightPartRoot := &TreeNode{}
		treeNode.Right = rightPartRoot
		createBTreeByOrderedArray(rightPartRoot, array, rightPartLeft, rightPartRight)
	}
}

func inorderTraversalBTree(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalBTree(root.Left)
	fmt.Printf(" %d", root.Val)
	inorderTraversalBTree(root.Right)
}

func main() {
	tests := [5][]int{
		{},                        //empty array
		{1},                       //one element
		{1, 2},                    //two element
		{1, 2, 3},                 //three element
		{1, 2, 3, 4, 5, 6, 7, 10}, //eight element
	}
	for _, test := range tests {
		root := CreateBTreeByOrderedArray(test)
		inorderTraversalBTree(root)
		fmt.Println()
	}
}
