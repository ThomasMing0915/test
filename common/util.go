package common

import "fmt"

//create single link list by variable parameters
func CreateSingleLinkListBySlice(slice ...interface{}) *LNode {
	//create head node
	headNode := NewLNode(nil)

	//create new node and add it to link list tail
	pNode := headNode
	for _, s := range slice {
		pNode.Next = NewLNode(s)
		pNode = pNode.Next
	}

	return headNode
}

func PrintSingleLinkList(info string, node *LNode) {
	if node == nil {
		return
	}

	fmt.Print(info, ":")

	//for node=node.next首次跳过头节点
	isFirstNodePrinted := false
	for node = node.Next; node != nil; node = node.Next {
		if !isFirstNodePrinted {
			fmt.Print(node.Data)
			isFirstNodePrinted = true
			continue
		}
		fmt.Print("->", node.Data)
	}
}
