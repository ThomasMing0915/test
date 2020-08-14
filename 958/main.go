package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isComplete bool
var flag bool

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isComplete = true
	flag = false //do not forget recover origin value
	dfs(root, depth(root), 1)
	return isComplete
}

func dfs(root *TreeNode, depth int, curDep int) {
	if root == nil {
		return
	}
	if curDep < depth-1 {
		if root.Left == nil || root.Right == nil {
			isComplete = false
			return
		}
	}
	if curDep+1 == depth {
		if root.Left == nil && root.Right != nil {
			isComplete = false
			return
		} else if root.Left != nil && root.Right == nil {
			if flag {
				isComplete = false
				return
			}
			flag = true
		} else if root.Left != nil && root.Right != nil {
			if flag {
				isComplete = false
			}
		} else { //why add this ?
			flag = true
			//    1
			//  2   3
			//     4 5
		}
	}
	dfs(root.Left, depth, curDep+1)
	dfs(root.Right, depth, curDep+1)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}

	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}

	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6
	fmt.Println(isCompleteTree(root))
}
