package main

import "fmt"

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid), len(grid))

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]), len(grid[i]))
	}

	//填充dp的第一列和第一行
	for j := 0; len(grid) > 0 && j < len(grid[0]); j++ {
		if j == 0 {
			dp[0][j] = grid[0][j]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}

	for i := 1; i < len(grid) && len(grid[0]) > 0; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	matrix := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(maxValue(matrix))
}
