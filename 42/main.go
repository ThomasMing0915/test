package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	// can use o(1) space, not use dp array?
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = math.MinInt32
	}
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
