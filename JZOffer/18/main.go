package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		next := head.Next
		head.Next = nil
		return next
	}

	p, q := head, head.Next
	for q != nil {
		if q.Val == val {
			break
		}
		p = p.Next
		q = q.Next
	}
	if q != nil {
		p.Next = q.Next
		q.Next = nil
	}

	return head
}
