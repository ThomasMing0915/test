package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	sourceNodeMap := make(map[*Node]int, 500)
	targetNodeMap := make(map[int]*Node, 500)

	node := head
	target := &Node{}
	travelNode := target

	pos := 1
	for node != nil {
		//fmt.Println(node.Val)
		sourceNodeMap[node] = pos
		travelNode.Next = &Node{
			Val:    node.Val,
			Next:   nil,
			Random: nil,
		}
		targetNodeMap[pos] = travelNode.Next
		travelNode = travelNode.Next
		node = node.Next
		pos++
	}

	travelNode = target.Next
	node = head
	for node != nil && travelNode != nil {
		travelNode.Random = targetNodeMap[sourceNodeMap[node.Random]]
		node = node.Next
		travelNode = travelNode.Next
	}

	return target.Next
}

func print(node *Node) {
	for node != nil {
		fmt.Printf("node:%d", node.Val)
		if node.Random != nil {
			fmt.Printf("   random:%d\n", node.Random.Val)
		} else {
			fmt.Printf("   random: \n")
		}
		node = node.Next
	}
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}

	node1.Next = node2
	node1.Random = node2
	node2.Random = node2

	node := copyRandomList(node1)
	print(node)
}
