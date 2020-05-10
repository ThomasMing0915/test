package main

import "fmt"

func singleNumber(nums []int) int {
	base := nums[0]
	for i := 1; i < len(nums); i++ {
		base ^= nums[i]
	}
	return base
}

func main() {
	tests := [3][]int{
		{1},
		{1, 2, 2},
		{1, 3, 1, 2, 2},
	}
	for _, test := range tests {
		fmt.Println(singleNumber(test))
	}
}
