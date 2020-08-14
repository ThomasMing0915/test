package main

import "fmt"

func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i-2 >= 0 && i-3 >= 0 {
			dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
		} else {
			if i < 2 {
				dp[i] = nums[i]
			} else {
				dp[i] = nums[2] + nums[0]
			}
		}
	}

	maxValue := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3}
	fmt.Println(rob(nums))
}
