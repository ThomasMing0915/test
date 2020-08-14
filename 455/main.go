package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	max := 0
	startPos := 0

	for i := 0; i < len(g); i++ {
		pos := find(startPos, g[i], s)
		if pos == -1 {
			break
		}
		max++
		startPos = pos + 1
	}

	return max
}

func find(startPos int, value int, s []int) int {
	for i := startPos; i < len(s); i++ {
		if s[i] >= value {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
