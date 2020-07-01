package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	//visit A
	if compare(A, B) {
		return true
	}
	//recursive A.Left  and recursive A.Right
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func compare(A *TreeNode, B *TreeNode) bool {
	if A == nil && B != nil {
		return false
	} else if A != nil && B == nil {
		return true
	} else if A == nil && B == nil {
		return true
	} else if A.Val != B.Val {
		return false
	}

	return compare(A.Left, B.Left) && compare(A.Right, B.Right)
}
