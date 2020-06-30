package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 {
		panic(k)
	}

	p, q := head, head
	i := 0

	for q != nil {
		i++
		q = q.Next

		if i == k {
			break
		}
	}

	if q == nil {
		if i == k {
			return p
		}
		panic("k is more than list length")
	}

	for q.Next != nil {
		q = q.Next
		p = p.Next
	}
	return p.Next
}
