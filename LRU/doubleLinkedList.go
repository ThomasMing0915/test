package main

import (
	"fmt"
)

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func NewNode(key, value int) *node {
	return &node{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func (n *node) String() string {
	return fmt.Sprintf("{%d:%d}", n.key, n.value)
}

/*---------------------------------------------------------*/

type DoubleLinkedList struct {
	capacity int
	head     *node
	tail     *node
	size     int
}

func NewDoubleLinkedList(cap int) *DoubleLinkedList {
	return &DoubleLinkedList{
		capacity: cap,
		head:     nil,
		tail:     nil,
		size:     0,
	}
}

//从头部添加元素
func (d *DoubleLinkedList) addHead(n *node) *node {
	if n == nil {
		return nil
	}

	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
	d.size += 1
	return n
}

//从尾部添加节点
func (d *DoubleLinkedList) addTail(n *node) *node {
	if n == nil {
		return nil
	}

	if d.tail == nil {
		d.head = n
		d.tail = n
	} else {
		d.tail.next = n
		n.prev = d.tail
		d.tail = n
	}
	d.size += 1
	return n
}

//任意节点删除
func (d *DoubleLinkedList) remove(n *node) *node {
	if d.tail == nil {
		return nil
	}
	//尾部节点
	if d.tail == n {
		d.deleteTail()
	} else if d.head == n {
		d.deleteHead()
	} else {
		n.next.prev = n.prev
		n.prev.next = n.next
		d.size--
	}

	return n
}

//从头部删除节点
func (d *DoubleLinkedList) deleteHead() *node {
	if d.head == nil {
		return nil
	}

	head := d.head

	if d.head.next != nil {
		tmp := d.head.next
		d.head.next = nil
		d.head = tmp
		d.head.prev = nil
	} else { //只有一个head节点
		d.head = nil
		d.tail = nil
	}
	d.size--
	return head
}

func (d *DoubleLinkedList) deleteTail() *node {
	if d.tail == nil {
		return nil
	}

	tail := d.tail

	if d.tail.prev != nil {
		tmp := d.tail.prev
		d.tail.prev = nil
		d.tail = tmp
		d.tail.next = nil
	} else {
		d.head = nil
		d.tail = nil
	}
	d.size--
	return tail
}

func (d *DoubleLinkedList) Pop() *node {
	return d.deleteHead()
}

func (d *DoubleLinkedList) Append(n *node) *node {
	return d.addTail(n)
}

func (d *DoubleLinkedList) AppendFront(n *node) *node {
	return d.addHead(n)
}

func (d *DoubleLinkedList) Remove(n *node) *node {
	return d.remove(n)
}

func (d *DoubleLinkedList) Print() {
	p := d.head
	line := ""
	for p != nil {
		line = line + fmt.Sprintf("%s", p)
		p = p.next
		if p != nil {
			line = line + "=>"
		}
	}
	fmt.Println(line)
}

func main() {
	l := NewDoubleLinkedList(10)
	nodes := make([]*node, 0, 10)
	for i := 0; i < 10; i++ {
		nodes = append(nodes, NewNode(i, i))
	}

	l.Append(nodes[0])
	l.Print()
	l.Append(nodes[1])
	l.Print()
	l.Pop()
	l.Print()
	l.Append(nodes[2])
	l.Print()
	l.AppendFront(nodes[3])
	l.Print()
	l.AppendFront(nodes[4])
	l.Print()
	l.Remove(nodes[2])
	l.Print()
}
