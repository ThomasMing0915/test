package main

import "fmt"

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			return nums[i]
		}
		m[nums[i]] = struct{}{}
	}

	//题目未定义能否走到这里，按原则是不能走到这里
	return -1
}

func main() {
	tests := [][]int{
		{2, 3, 1, 0, 2, 5, 3},
		{0, 1, 2, 3, 3},
		{0, 0, 0},
		{0, 1},
	}
	for _, test := range tests {
		fmt.Println(findRepeatNumber(test))
	}
}
