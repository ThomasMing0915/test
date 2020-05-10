package main

import "fmt"

func main() {
	linkMode()
}

func linkMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ls := &linkStack{node: NewLinkNode(-1)}

	ls.Push(1)
	ls.Push(2)

	fmt.Println("栈的大小为:", ls.Size())
	fmt.Println("栈顶元素为:", ls.Top())

	ls.Pop()
	fmt.Println("栈的大小为:", ls.Size())
	fmt.Println("栈顶元素为:", ls.Top())

	ls.Pop()
	fmt.Println("栈的大小为:", ls.Size())
	fmt.Println("栈顶元素为:", ls.Top())
}
