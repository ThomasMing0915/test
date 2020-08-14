package main

import "fmt"

//单链表节点定义
type node struct {
	value int   //节点的值
	next  *node //指向下一个节点的指针
}

func NewNode(value int) *node {
	n := &node{
		value: value,
		next:  nil,
	}
	return n
}

//根据int切片创建一个带头节点的单链表
func CreateSingleLinkList(ints []int) *node {
	head := NewNode(0) //头节点
	pnode := head      //移动指针

	//尾插法创建单链表
	for _, v := range ints {
		newnode := NewNode(v)
		pnode.next = newnode
		pnode = pnode.next
	}

	return head
}

//原地逆序单链表
func ReverseSingleLinkList(n *node) *node {
	if n == nil {
		return nil
	}
	head := n             //头节点,value不具有实际意义
	pReversedNode := head //已逆序部分的头节点
	if head.next == nil {
		return head
	}
	pHeadNext := head.next

	pCurrentNode := head.next //未逆序部分的起始节点
	for pCurrentNode != nil {
		pNextNode := pCurrentNode.next //未逆序部分起始节点的下一个节点
		//将pCurrentNode加入到逆序链表中
		pCurrentNode.next = pReversedNode
		pReversedNode = pCurrentNode
		pCurrentNode = pNextNode
	}

	head = nil
	pHeadNext.next = nil
	newHead := NewNode(0)
	newHead.next = pReversedNode
	return newHead
}

/*--------------------------------------------------------------------------*/
func ReverseSingleLinkListByRecursive(n *node) *node {
	//递归终止条件
	if n == nil || n.next == nil {
		return n
	}

	//递归处理完成从n.next位置节点的链表的逆序，并返回新的头节点
	newNode := ReverseSingleLinkListByRecursive(n.next)
	pNode := newNode
	//pNode依次从头遍历到尾部
	for pNode.next != nil {
		pNode = pNode.next
	}
	pNode.next = n
	n.next = nil

	return newNode
}

func PrintSingleLinkList(n *node) {
	if n != nil {
		fmt.Print(n.value)
		n = n.next
	}
	for n != nil {
		fmt.Print("->", n.value)
		n = n.next
	}
	fmt.Println()
}
