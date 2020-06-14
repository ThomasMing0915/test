package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	mergeListHead := &ListNode{}
	p := mergeListHead
	p0 := lists[0]
	for p0 != nil {
		p.Next = &ListNode{Val: p0.Val}
		p = p.Next
		p0 = p0.Next
	}
	for i := 1; i < len(lists); i++ {
		mergeListHead = mergeTwoLists(mergeListHead, lists[i])
	}
	return mergeListHead.Next
}

func mergeTwoLists(listA, listB *ListNode) *ListNode {
	head := listA //哨兵节点
	p := listA

	////merge B to A
	listA = listA.Next
	for listA != nil && listB != nil {
		if listA.Val <= listB.Val {
			p = p.Next
			listA = listA.Next
		} else {
			tmpListB := listB.Next
			//tmpListA:=listA.Next

			listB.Next = p.Next
			p.Next = listB

			p = p.Next
			//listA = listA.Next

			//listA=tmpListA
			listB = tmpListB
		}
	}

	if listB != nil {
		p.Next = listB
	}

	return head
}

func main() {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 1}
	list2.Next.Next = &ListNode{Val: 1}

	list3 := &ListNode{Val: -2}
	list3.Next = &ListNode{Val: 6}

	mergelist := mergeKLists([]*ListNode{list1, list2, list3})
	for mergelist != nil {
		fmt.Println(mergelist.Val)
		mergelist = mergelist.Next
	}

}
