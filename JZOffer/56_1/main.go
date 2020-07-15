package main

import "fmt"

func singleNumbers(nums []int) []int {
	//1. 全部数异或res
	var res int
	for _, num := range nums {
		res = res ^ num
	}

	//2. 找到res中一个不为0的bit位  ... 0000 0100 0000
	var dim = 1
	for {
		if dim&res != 0 {
			break
		}
		dim = dim << 1
	}

	//3.将nums根据2中不为1的bit是0还是1分成两组 分别进行异或
	var a, b int
	for _, num := range nums {
		if dim&num != 0 {
			a = a ^ num
		} else {
			b = b ^ num
		}
	}
	return []int{a, b}
}

func main() {
	nums := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(nums))
}
