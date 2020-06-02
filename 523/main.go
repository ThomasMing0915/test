package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			//sumij=sum[i]+sum[i+1]+...+sum[j]
			sumij := prefixSum[j] - prefixSum[i] + nums[i]
			// k may be is 0
			if sumij == k || (k != 0 && sumij%k == 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{23, 2, 4, 6, 7, -1, 1}
	fmt.Println(checkSubarraySum(nums, 0))
}
