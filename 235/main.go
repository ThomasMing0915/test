package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if oneIsAnotherParent(p, q) {
		return p
	}
	if oneIsAnotherParent(q, p) {
		return q
	}

	for {
		pParent, lp := findNodeDirectParent(root, p, 0)
		qParent, lq := findNodeDirectParent(root, q, 0)
		if pParent == qParent {
			return pParent
		}
		if lp > lq {
			p = pParent
		} else if lp < lq {
			q = qParent
		} else {
			p = pParent
			q = qParent
		}
	}
}

func oneIsAnotherParent(q, p *TreeNode) bool {
	if q == p {
		return true
	}

	if q.Left != nil {
		if oneIsAnotherParent(q.Left, p) {
			return true
		}
	}

	if q.Right != nil {
		if oneIsAnotherParent(q.Right, p) {
			return true
		}
	}

	return false
}

func findNodeDirectParent(root, q *TreeNode, level int) (*TreeNode, int) {
	if root.Left == q || root.Right == q || root == q {
		return root, level
	}

	if root.Left != nil {
		parent, l := findNodeDirectParent(root.Left, q, level+1)
		if parent != nil {
			return parent, l
		}
	}

	if root.Right != nil {
		parent, l := findNodeDirectParent(root.Right, q, level+1)
		if parent != nil {
			return parent, l
		}
	}

	return nil, 0
}
