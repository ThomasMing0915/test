package main

import (
	"LeetCode/common"
	"fmt"
)

func main() {
	head := common.CreateSingleLinkListBySlice(1, 3, 1, 5, 5, 7)
	head = RemoveDupItem(head)
	common.PrintSingleLinkList("remove dup node", head)

	fmt.Println("\n通过递归删除单链表中重复的元素")
	head = common.CreateSingleLinkListBySlice(1, 3, 1, 5, 5, 7)
	head = RemoveDupItemByRecursive(head)
	common.PrintSingleLinkList("remove dup item by recursive", head)
}

func RemoveDupItem(node *common.LNode) *common.LNode {
	if node == nil || node.Next == nil {
		return node
	}
	//pSortedNodeStart和End分别表示已验证过区间的起始和终止节点
	pSortedNodeStart := node
	pSortedNodeEnd := node

	for {
		//pCurrentNode表示待验证的节点
		pCurrentNode := pSortedNodeEnd.Next
		if pCurrentNode == nil {
			break
		}

		find := false
		for pNode := pSortedNodeStart; pNode != pSortedNodeEnd.Next; pNode = pNode.Next {
			iCurrentData, ok1 := pCurrentNode.Data.(int)
			iNodeData, ok2 := pNode.Data.(int)

			if ok1 && ok2 {
				if iCurrentData == iNodeData {
					find = true
					break
				}
			}
		}
		if find { //找到了 跳过当前元素
			pSortedNodeEnd.Next = pCurrentNode.Next
		} else {
			pSortedNodeEnd = pCurrentNode
		}
	}
	return pSortedNodeStart
}

func RemoveDupItemByRecursive(node *common.LNode) *common.LNode {
	return removeDupItemByRecursive(node)
}

func removeDupItemByRecursive(node *common.LNode) *common.LNode {
	//remove sub link list dup item
	//递归终止条件
	if node == nil || node.Next == nil {
		return node
	}

	//先处理子链表重复元素
	subHead := removeDupItemByRecursive(node.Next)

	node.Next = subHead
	pNode := node

	for pCurrent := pNode.Next; pCurrent != nil; pCurrent = pCurrent.Next {
		iCurrentData, ok1 := pCurrent.Data.(int)
		iNodeData, ok2 := node.Data.(int)
		if ok1 && ok2 {
			if iCurrentData == iNodeData {
				pNode.Next = pCurrent.Next
				break //与子链表中最多存在一个相同元素
			} else {
				pNode = pCurrent
			}
		}
	}
	return node
}
