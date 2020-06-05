package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var lastVisitNode *TreeNode
var maxAppearNums map[int][]int //key:appear numbs  value:what num appears max count
var appearNums int

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	maxAppearNums = make(map[int][]int)
	appearNums = 0
	lastVisitNode = nil

	dfs(root)

	if len(maxAppearNums) == 0 {
		maxAppearNums[appearNums] = []int{lastVisitNode.Val}
	} else {
		var k int
		var v []int
		for k, v = range maxAppearNums {
			if k == appearNums {
				vv := make([]int, 0, len(v)+1)
				vv = append(vv, v...)
				vv = append(vv, lastVisitNode.Val)
				maxAppearNums[k] = vv
			} else if k < appearNums {
				//maxAppearNums[appearNums] = []int{lastVisitNode.Val}
				delete(maxAppearNums, k)
			}
		}
		if k < appearNums {
			maxAppearNums[appearNums] = []int{lastVisitNode.Val}
		}
	}

	for _, v := range maxAppearNums {
		return v
	}
	return nil
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	if lastVisitNode != nil {
		if lastVisitNode.Val == node.Val {
			appearNums++
		} else {
			if len(maxAppearNums) == 0 {
				maxAppearNums[appearNums] = []int{lastVisitNode.Val}
			} else {
				var k int
				var v []int
				for k, v = range maxAppearNums {
					if k == appearNums {
						vv := make([]int, 0, len(v)+1)
						vv = append(vv, v...)
						vv = append(vv, lastVisitNode.Val)
						maxAppearNums[k] = vv
					} else if k < appearNums {
						delete(maxAppearNums, k)
					}
				}
				if k < appearNums {
					maxAppearNums[appearNums] = []int{lastVisitNode.Val}
				}
			}
			appearNums = 1
		}
	} else {
		appearNums = 1
	}
	lastVisitNode = node
	dfs(node.Right)
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: -2}
	node3 := &TreeNode{Val: 1}
	node11 := &TreeNode{Val: 1}
	node3x := &TreeNode{Val: 3}
	node31 := &TreeNode{Val: 3}
	node32 := &TreeNode{Val: 3}

	root.Left = node1
	root.Right = node11
	node1.Left = node2
	node1.Right = node3
	node11.Right = node3x
	node3x.Left = node31
	node3x.Right = node32

	fmt.Println(findMode(root))
}
