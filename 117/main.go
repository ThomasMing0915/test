package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	connectInner2(root)
	return root
}

//wrong answer why?
//func connectInner(node *Node) {
//	if node == nil {
//		return
//	}
//
//	//leaf node
//	if node.Left == nil && node.Right == nil {
//		return
//	}
//
//	if node.Left != nil && node.Right == nil {
//		if node.Next != nil {
//			if node.Next.Left != nil {
//				node.Left.Next = node.Next.Left
//			} else if node.Next.Right != nil {
//				node.Left.Next = node.Next.Right
//			}
//		}
//	}
//
//	if node.Left == nil && node.Right != nil {
//		if node.Next != nil {
//			if node.Next.Left != nil {
//				node.Right.Next = node.Next.Left
//			} else if node.Next.Right != nil {
//				node.Right.Next = node.Next.Right
//			}
//		}
//	}
//
//	if node.Left != nil && node.Right != nil {
//		node.Left.Next = node.Right
//		if node.Next != nil {
//			if node.Next.Left != nil {
//				node.Right.Next = node.Next.Left
//			} else if node.Next.Right != nil {
//				node.Right.Next = node.Next.Right
//			}
//		}
//	}
//	connectInner(node.Left)
//	connectInner(node.Right)
//}

//wrong answer why?
func connectInner2(node *Node) {
	if node == nil {
		return
	}

	//leaf node
	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Left != nil && node.Right == nil {
		p := node
		for p.Next != nil {
			if p.Next.Left != nil {
				node.Left.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				node.Left.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}

	if node.Left == nil && node.Right != nil {
		p := node
		for p.Next != nil {
			if p.Next.Left != nil {
				node.Right.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				node.Right.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}

	if node.Left != nil && node.Right != nil {
		node.Left.Next = node.Right
		p := node
		for p.Next != nil {
			if p.Next.Left != nil {
				node.Right.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				node.Right.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}
	connectInner2(node.Left) //先left后right 还是先right后left
	connectInner2(node.Right)
}

func connectInner3(node *Node) {
	if node == nil {
		return
	}

	//leaf node
	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Left != nil && node.Right == nil {
		p := node
		for p.Next != nil {
			if p.Next.Left != nil {
				node.Left.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				node.Left.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}

	if node.Left == nil && node.Right != nil {
		p := node
		for p.Next != nil {
			if p.Next.Left != nil {
				node.Right.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				node.Right.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}

	if node.Left != nil && node.Right != nil {
		node.Left.Next = node.Right
		p := node
		for p.Next != nil {
			if p.Next.Left != nil {
				node.Right.Next = p.Next.Left
				break
			} else if p.Next.Right != nil {
				node.Right.Next = p.Next.Right
				break
			}
			p = p.Next
		}
	}
	connectInner3(node.Right)
	connectInner3(node.Left)
}

func travelNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("\ncurNode val :%d ", node.Val)
	if node.Left != nil {
		fmt.Printf(" leftNode val :%d ", node.Left.Val)
	}
	if node.Right != nil {
		fmt.Printf(" rightNode val :%d ", node.Right.Val)
	}
	if node.Next != nil {
		fmt.Printf(" nextNode val :%d ", node.Next.Val)
	}
	travelNode(node.Left)
	travelNode(node.Right)
}

func main() {
	root := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node4.Left = node7
	node3.Right = node6
	node6.Left = node8

	connect(root)
	travelNode(root)
}
