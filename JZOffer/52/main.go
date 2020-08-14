package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listACount := listNodeCount(headA)
	listBCount := listNodeCount(headB)

	if listACount > listBCount {
		diff := listACount - listBCount
		for i := 0; i < diff; i++ {
			headA = headA.Next
		}
	} else {
		diff := listBCount - listACount
		for i := 0; i < diff; i++ {
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func listNodeCount(head *ListNode) int {
	var count int
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
