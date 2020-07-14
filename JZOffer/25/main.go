package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var mergeHead ListNode
	p := &mergeHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			p.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		p = p.Next
	}

	for l1 != nil {
		p.Next = &ListNode{
			Val: l1.Val,
		}
		l1 = l1.Next
		p = p.Next
	}

	for l2 != nil {
		p.Next = &ListNode{
			Val: l2.Val,
		}
		l2 = l2.Next
		p = p.Next
	}
	return mergeHead.Next
}
