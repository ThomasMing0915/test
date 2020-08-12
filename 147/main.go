package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func insertionSortList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}

	fackHead:=&ListNode{}
	fackHead.Next=head

	p:=head
	pn:=p.Next

	for pn!=nil{
		if pn.Val<p.Val{
			tmp:=fackHead
			for tmp!=nil && tmp.Next!=nil && tmp.Next.Val<=pn.Val{
				tmp=tmp.Next
			}
			tail:=pn.Next
			pn.Next=tmp.Next
			tmp.Next=pn

			p.Next=tail
			pn=tail

			continue
		}
		p=pn
		pn=pn.Next
	}

	return fackHead.Next
}


func main(){
	head:=&ListNode{Val: 4,Next: nil}
	node2:=&ListNode{Val: 2,Next: nil}
	node1:=&ListNode{Val: 1,Next: nil}
	node3:=&ListNode{Val: 3,Next: nil}

	head.Next=node2
	node2.Next=node1
	node1.Next=node3

	newHead:=insertionSortList(head)
	for newHead!=nil{
		fmt.Println(newHead.Val)
		newHead=newHead.Next
	}
}