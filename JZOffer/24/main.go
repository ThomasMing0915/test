package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil { //其实这个判断可以合并到下面的for中,即这里的if可以去掉
		return head
	}

	p := head
	q := head.Next

	for q != nil {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}
	head.Next = nil

	return p
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("->%d", head.Val)
		head = head.Next
	}
}

func reverseListByRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead, _ := recursive(head)
	return newHead
}

func recursive(node *ListNode) (*ListNode, *ListNode) {
	if node.Next == nil {
		return node, node
	}
	head, tail := recursive(node.Next)
	tail.Next = node
	node.Next = nil
	tail = node

	return head, tail
}

//上面的recursive要函数要返回两个参数，能不能直接返回一个? 第二个返回参数tail能不能省略掉
//tail与node.Next是啥关系，其实node.Next就是tail,不明白可以画个图,所以上面的递归函数可以
//简化
func recursive2(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	head := recursive2(node.Next)
	node.Next.Next = node
	node.Next = nil

	return head
}

func main() {
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	head = reverseList(head)
	print(head)

	head = reverseListByRecursive(head)
	print(head)

}
