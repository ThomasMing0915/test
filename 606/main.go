package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gstr []string

func tree2str(t *TreeNode) string {
	gstr = make([]string, 0, 32)

	dfs(t)

	var b strings.Builder
	for _, str := range gstr {
		b.WriteString(str)
	}
	return b.String()

}

func dfs(t *TreeNode) {
	if t == nil {
		return
	}
	gstr = append(gstr, strconv.Itoa(t.Val))

	if t.Left == nil && t.Right == nil {
		return
	}

	if t.Left != nil && t.Right != nil {
		gstr = append(gstr, "(")
		dfs(t.Left)
		gstr = append(gstr, ")")

		gstr = append(gstr, "(")
		dfs(t.Right)
		gstr = append(gstr, ")")
	} else if t.Left != nil && t.Right == nil {
		gstr = append(gstr, "(")
		dfs(t.Left)
		gstr = append(gstr, ")")
	} else {
		gstr = append(gstr, "(")
		gstr = append(gstr, ")")

		gstr = append(gstr, "(")
		dfs(t.Right)
		gstr = append(gstr, ")")
	}

}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}

	root.Left = node2
	root.Right = node3
	node2.Right = node4

	fmt.Println(tree2str(root))
}
