package main

type linkNode struct {
	val  int
	next *linkNode
}

func NewLinkNode(v int) *linkNode {
	return &linkNode{val: v}
}

type linkStack struct {
	node *linkNode
	size int
}

func (l *linkStack) IsEmpty() bool {
	return l.size == 0
}

func (l *linkStack) Size() int {
	return l.size
}

func (l *linkStack) Push(v int) {
	ln := NewLinkNode(v)
	ln.next = l.node.next
	l.node.next = ln
	l.size += 1
}

func (l *linkStack) Pop() int {
	if l.IsEmpty() {
		panic("link stack is empty, can't pop")
	}
	tNode := l.node.next
	if tNode == nil {
		panic("link stack is empty, can't pop")
	}

	l.node.next = tNode.next
	tNode.next = nil //将tNsode从链表移除
	l.size -= 1

	return tNode.val
}

func (l *linkStack) Top() int {
	if l.IsEmpty() {
		panic("link stack is empty, can't top")
	}

	if l.node.next == nil {
		panic("link stack is empty, can't top")
	}

	return l.node.next.val
}
