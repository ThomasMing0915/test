package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil && t2 != nil {
		return t2
	}
	if t1 != nil && t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	mergeTreesInner(t1, t2)
	return t1
}

func mergeTreesInner(t1 *TreeNode, t2 *TreeNode) {
	if t1 == nil && t2 == nil {
		return
	}
	leftRecursive := true
	rightRecursive := true
	if t1.Left == nil && t2.Left != nil {
		t1.Left = t2.Left
		leftRecursive = false
	} else if t1.Left != nil && t2.Left == nil {
		leftRecursive = false
	} else if t1.Left == nil && t2.Left == nil {
		leftRecursive = false
	} else if t1.Left != nil && t2.Left != nil {
		t1.Left.Val = t1.Left.Val + t2.Left.Val
	}

	//why blew use else if ? if use if what happen
	if t1.Right == nil && t2.Right != nil {
		t1.Right = t2.Right
		rightRecursive = false
	} else if t1.Right != nil && t2.Right == nil {
		rightRecursive = false
	} else if t1.Right == nil && t2.Right == nil {
		rightRecursive = false
	} else if t1.Right != nil && t2.Right != nil {
		t1.Right.Val = t1.Right.Val + t2.Right.Val
	}
	if leftRecursive {
		mergeTreesInner(t1.Left, t2.Left)
	}
	if rightRecursive {
		mergeTreesInner(t1.Right, t2.Right)
	}
}
