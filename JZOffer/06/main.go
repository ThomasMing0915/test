package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	arr := reversePrint(head.Next)
	arr = append(arr, head.Val)
	return arr
}
