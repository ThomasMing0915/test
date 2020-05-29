package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 != nil && root2 != nil {
		if root1.Val != root2.Val {
			return false
		}
		if root1.Left != nil && root1.Right != nil {
			if root2.Left != nil && root2.Right != nil {
				if root1.Left.Val == root2.Left.Val && root1.Right.Val == root2.Right.Val {
					if root1.Left.Val != root1.Right.Val {
						return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
					} else {
						return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
							(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
					}
				} else if root1.Left.Val == root2.Right.Val && root1.Right.Val == root2.Left.Val {
					return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
				}
			} else {
				return false
			}
		} else if root1.Left == nil && root1.Right == nil {
			if root2.Left == nil && root2.Right == nil {
				return true
			} else {
				return false
			}
		} else {
			if root2.Left == nil && root2.Right == nil {
				return false
			} else if root2.Left != nil && root2.Right != nil {
				return false
			} else {
				if root1.Left != nil && root2.Left != nil {
					if root1.Left.Val == root2.Left.Val {
						return flipEquiv(root1.Left, root2.Left)
					} else {
						return false
					}
				} else if root1.Left != nil && root2.Right != nil {
					if root1.Left.Val == root2.Right.Val {
						return flipEquiv(root1.Left, root2.Right)
					} else {
						return false
					}
				} else if root1.Right != nil && root2.Left != nil {
					if root1.Right.Val == root2.Left.Val {
						return flipEquiv(root1.Right, root2.Left)
					} else {
						return false
					}
				} else if root1.Right != nil && root2.Right != nil {
					if root1.Right.Val == root2.Right.Val {
						return flipEquiv(root1.Right, root2.Right)
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}
