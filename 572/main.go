package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if equal(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equal(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s != nil && t != nil && s.Val == t.Val {
		return equal(s.Left, t.Left) && equal(s.Right, t.Right)
	}
	return false
}

func main() {
	s := &TreeNode{Val: 3}
	nodes4 := &TreeNode{Val: 4}
	nodes5 := &TreeNode{Val: 5}
	nodes1 := &TreeNode{Val: 1}
	nodes2 := &TreeNode{Val: 2}
	s.Left = nodes4
	s.Right = nodes5
	nodes4.Left = nodes1
	nodes4.Right = nodes2

	t := &TreeNode{Val: 4}
	nodet1 := &TreeNode{Val: 1}
	nodet2 := &TreeNode{Val: 2}
	t.Left = nodet1
	t.Right = nodet2

	fmt.Println(isSubtree(s, t))
}
