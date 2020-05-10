package main

import (
	"LeetCode/common"
	"fmt"
)

func main() {
	head := common.CreateSingleLinkListBySlice(1, 2, 3, 4, 5)
	common.PrintSingleLinkList("重新排序前 ", head)
	ReSortSingleLinkList(head.Next)
	common.PrintSingleLinkList(" 重新排序后 ", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice(1)
	common.PrintSingleLinkList("重新排序前 ", head)
	ReSortSingleLinkList(head.Next)
	common.PrintSingleLinkList(" 重新排序后 ", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice(1, 2)
	common.PrintSingleLinkList("重新排序前 ", head)
	ReSortSingleLinkList(head.Next)
	common.PrintSingleLinkList(" 重新排序后 ", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice(1, 2, 3)
	common.PrintSingleLinkList("重新排序前 ", head)
	ReSortSingleLinkList(head.Next)
	common.PrintSingleLinkList(" 重新排序后 ", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice()
	common.PrintSingleLinkList("重新排序前 ", head)
	ReSortSingleLinkList(head.Next)
	common.PrintSingleLinkList(" 重新排序后 ", head)

	fmt.Println("\n通过递归方法重新排序")
	head = common.CreateSingleLinkListBySlice(1, 2, 3, 4, 5)
	common.PrintSingleLinkList("重新排序前", head)
	ResortSingLinkByRecursive(head)
	common.PrintSingleLinkList("重新排序后", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice(1)
	common.PrintSingleLinkList("重新排序前", head)
	ResortSingLinkByRecursive(head)
	common.PrintSingleLinkList("重新排序后", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice(1, 2)
	common.PrintSingleLinkList("重新排序前", head)
	ResortSingLinkByRecursive(head)
	common.PrintSingleLinkList("重新排序后", head)

	fmt.Println()
	head = common.CreateSingleLinkListBySlice(1, 2, 3)
	common.PrintSingleLinkList("重新排序前", head)
	ResortSingLinkByRecursive(head)
	common.PrintSingleLinkList("重新排序后", head)
}

func ReSortSingleLinkList(node *common.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	//pNextNode表示还未进行重排链表的开始节点
	pNextNode := node.Next
	//pCurrentNode表示已经排好序的链表的尾节点
	pCurrentNode := node
	pNode := pNextNode
	pPreNode := node

	for {
		if pNode == nil || pNode.Next == nil {
			break
		}

		//pNextNode停留在尾节点
		for pNode.Next != nil {
			pNode = pNode.Next
			pPreNode = pPreNode.Next
		}
		pCurrentNode.Next = pNode
		pNode.Next = pNextNode
		pPreNode.Next = nil

		pCurrentNode = pNextNode
		pNextNode = pCurrentNode.Next
		pNode = pNextNode
		pPreNode = pCurrentNode
	}
}

func ResortSingLinkByRecursive(node *common.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	//node.Next跳过头节点
	resortSingleLink(node.Next)
}

func resortSingleLink(node *common.LNode) {
	//递归结束条件，只有一个节点或没有节点时
	if node == nil || node.Next == nil || node.Next.Next == nil {
		return
	}

	pNode := node
	for pNode != nil && pNode.Next != nil && pNode.Next.Next != nil {
		pNode = pNode.Next
	}

	pTail := pNode.Next
	pTail.Next = node.Next

	pNode.Next = nil
	node.Next = pTail

	resortSingleLink(pTail.Next)
}
