package main

import "fmt"

func main() {
	a1 := [3][2]int{}

	fmt.Println("a1", len(a1), len(a1[0]))

	//a2 动态初始化二维数组
	a2 := [len(a1)][len(a1[0])]int{}

	fmt.Println("a2", a2)

	fmt.Println("cap a2", cap(a2))

	////////
	var s []int
	s = append(s, 1)

}
