package main

type CQueue struct {
	stackMain Stack
	stackTmp  Stack
}

type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{
		stack: make([]int, 0, 10000),
	}
}

func (s *Stack) Push(v int) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func Constructor() CQueue {
	return CQueue{
		stackMain: NewStack(),
		stackTmp:  NewStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackMain.Push(value)
}

func (this *CQueue) DeleteHead() int {
	for {
		top := this.stackMain.Pop()
		if top == -1 {
			break
		}

		this.stackTmp.Push(top)
	}

	head := this.stackTmp.Pop()

	for {
		top := this.stackTmp.Pop()
		if top == -1 {
			break
		}

		this.stackMain.Push(top)
	}

	return head
}
