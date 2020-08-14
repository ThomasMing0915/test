package main

import "fmt"

func main() {
	fmt.Println("第一组测试")
	head := CreateSingleLinkList(nil)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead := ReverseSingleLinkList(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第二组测试")
	slice := []int{1}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkList(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第三组测试")
	slice = []int{1, 2}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkList(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第四组测试")
	slice = []int{1, 2, 3}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkList(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第五组测试")
	slice = []int{1, 2, 3, 4, 5, 6}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkList(head)
	PrintSingleLinkList(reverseHead)

	//通过递归方法
	fmt.Println("\n通过递归方法实现单链表逆序")
	fmt.Println("第一组测试")
	head = CreateSingleLinkList(nil)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkListByRecursive(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第二组测试")
	slice = []int{1}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkListByRecursive(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第三组测试")
	slice = []int{1, 2}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkListByRecursive(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第四组测试")
	slice = []int{1, 2, 3}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkListByRecursive(head)
	PrintSingleLinkList(reverseHead)

	fmt.Println("第五组测试")
	slice = []int{1, 2, 3, 4, 5, 6}
	head = CreateSingleLinkList(slice)
	PrintSingleLinkList(head)
	//fmt.Println()
	reverseHead = ReverseSingleLinkListByRecursive(head)
	PrintSingleLinkList(reverseHead)
}
