package main

import (
	"fmt"
)

//使用双链表实现LRU算法

type Data struct {
	key   int
	value int
}

type Node struct {
	data Data
	next *Node
	pre  *Node
}

type LRU struct {
	count int
	size  int
	head  *Node
}

func NewLRU() *LRU {
	return &LRU{
		count: 0,
		size:  5,
		head:  nil,
	}
}

func (l *LRU) Put(data Data) {
	node := &Node{
		data: data,
		next: nil,
		pre:  nil,
	}
	if l.IsFull() {
		//淘汰最末尾的节点
		node.next = l.head
		node.pre = l.head.pre.pre
		l.head.pre.pre.next = node
		l.head.pre = node
	} else if l.IsEmpty() {
		node.next = node
		node.pre = node

		l.head = node

		l.count++
	} else {
		node.next = l.head
		node.pre = l.head.pre

		l.head.pre.next = node
		l.head.pre = node
		l.count++
	}
}

func (l *LRU) Get(key int) int {
	node := l.head
	for i := 0; i < l.count; i++ {
		if node.data.key != key {
			node = node.next
			continue
		}

	}
	return -1
}

func (l *LRU) IsEmpty() bool {
	return l.count == 0
}

func (l *LRU) IsFull() bool {
	return l.count == l.size
}

func (l *LRU) Print() {
	fmt.Printf("lru count:%d\n", l.count)
	node := l.head
	for i := 0; i < l.count; i++ {
		fmt.Println(node.data.key, node.data.value)
		node = node.next
	}
}

//func main() {
//	lru := NewLRU()
//	lru.Print()
//	lru.Put(Data{key: 1, value: 11})
//	lru.Print()
//	lru.Put(Data{key: 2, value: 22})
//	lru.Put(Data{key: 3, value: 33})
//	lru.Print()
//	lru.Put(Data{key: 4, value: 44})
//	lru.Put(Data{key: 5, value: 55})
//	lru.Print()
//	lru.Put(Data{key: 6, value: 66})
//	lru.Print()
//}
