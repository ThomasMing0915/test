package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) []int {
	i := firstIncrease(nums)
	if i == -1 { //nums已全部是降序，没有比他更大的全排列了
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return nums
	}

	for j := len(nums) - 1; j >= i; j-- {
		if nums[j] > nums[i-1] {
			nums[j], nums[i-1] = nums[i-1], nums[j]
			sort.Slice(nums[i:], func(a, b int) bool {
				return nums[i+a] < nums[i+b] //!!! why use nums[i+a] nums[i+b] not use nums[a],nums[b]
			})
			return nums
		}
	}
	return nil
}

func firstIncrease(nums []int) int {
	//find A[i]>A[i-1]
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			return i
		}
	}
	//not found
	return -1
}

func main() {
	nums := []int{2, 3, 1}
	nums = nextPermutation(nums)
	fmt.Println(nums)
}
