package main

import "fmt"

func main() {
	SliceModel()
}

func SliceModel() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("用切片数组构建栈")
	ss := &sliceStack{arr: make([]int, 0)}

	ss.Push(1)
	ss.Push(2)
	fmt.Println("栈顶元素为:", ss.Top())
	fmt.Println("栈大小为:", ss.size)

	ss.Pop()
	fmt.Println("出栈成功:", ss.size)
	ss.Pop()

}
