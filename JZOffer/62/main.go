package main

import "fmt"

//用循环单链表来做
type LinkedList struct {
	Val  int
	Next *LinkedList
}

func newNode(n int) *LinkedList {
	return &LinkedList{
		Val: n,
	}
}

func createList(n int) *LinkedList {
	if n <= 0 {
		return nil
	}

	head := newNode(0)
	p := head

	for i := 1; i < n; i++ {
		p.Next = newNode(i)
		p = p.Next
	}

	p.Next = head

	return p
}

func lastRemaining(n int, m int) int {
	node := createList(n)

	pos := 0
	for node.Next != node {
		if pos == m-1 {
			node.Next = node.Next.Next
			pos = 0
			continue
		}
		node = node.Next
		pos++
		//fmt.Println(node.Val)
	}

	return node.Val
}

func main() {
	for _, num := range [][2]int{{5, 3}, {10, 17}, {9, 13} /* {70866, 116922}*/} {
		fmt.Println(lastRemaining(num[0], num[1]))
	}

}
