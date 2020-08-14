package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	return dfs(postorder, 0, len(postorder)-1)
}

func dfs(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}

	rootIndex := right

	var i, j int
	for i = left; i < rootIndex; i++ {
		if postorder[i] > postorder[rootIndex] {
			break
		}
	}
	for j = i; j < rootIndex; j++ {
		if postorder[j] < postorder[rootIndex] {
			return false
		}
	}

	return dfs(postorder, left, i-1) && dfs(postorder, i+1, right)
}

func main() {
	for _, arr := range [][]int{{1, 6, 2, 2, 5}, {1, 3, 2, 6, 5}} {
		fmt.Println(verifyPostorder(arr))
	}
}
