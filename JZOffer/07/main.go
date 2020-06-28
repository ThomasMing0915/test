package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{}

	dfs(preorder, inorder, root)
	return root
}

func dfs(preorder []int, inorder []int, node *TreeNode) {
	if len(preorder) == 0 || len(inorder) == 0 {
		return
	}

	node.Val = preorder[0]

	i := index(inorder, node.Val)

	if i > 0 {
		left := &TreeNode{}
		node.Left = left
		dfs(preorder[1:i+1], inorder[:i], left)
	}

	if len(inorder[i+1:]) > 0 {
		right := &TreeNode{}
		node.Right = right
		dfs(preorder[i+1:], inorder[i+1:], right)
	}

}

func index(order []int, v int) int {
	for i := 0; i < len(order); i++ {
		if order[i] == v {
			return i
		}
	}
	return -1
}
