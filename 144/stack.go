package main

type StackItem struct {
	item interface{}
	next *StackItem
}

type Stack struct {
	sp    *StackItem
	depth int64
}

func New() *Stack {
	var stack *Stack = new(Stack)

	stack.depth = 0
	return stack
}

func (stack *Stack) Push(item interface{}) {
	stack.sp = &StackItem{item: item, next: stack.sp}
	stack.depth++
}

func (stack *Stack) Pop() interface{} {
	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}
	return nil
}

func (stack *Stack) Peek() interface{} {
	if stack.depth > 0 {
		return stack.sp.item
	}
	return nil
}

func (stack *Stack) IsEmpty() bool {
	return stack.depth == 0
}
