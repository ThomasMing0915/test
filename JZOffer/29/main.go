package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	res := make([]int, 0, len(matrix)*len(matrix[0]))

	rowStart := 0
	rowEnd := len(matrix) - 1

	colStart := 0
	colEnd := len(matrix[0]) - 1

	for rowStart <= rowEnd {
		if rowStart == rowEnd {
			for i := colStart; i <= colEnd; i++ {
				res = append(res, matrix[rowStart][i])
			}
			break
		} else if colStart == colEnd {
			//fmt.Println(111, rowStart, rowEnd, colEnd)
			for i := rowStart; i <= rowEnd; i++ {
				res = append(res, matrix[i][colEnd])
			}
			break
		} else if colStart > colEnd {
			break
		} else {
			for i := colStart; i <= colEnd; i++ {
				res = append(res, matrix[rowStart][i])
			}
			for i := rowStart + 1; i < rowEnd; i++ {
				res = append(res, matrix[i][colEnd])
			}
			for i := colEnd; i >= colStart; i-- {
				res = append(res, matrix[rowEnd][i])
			}
			for i := rowEnd - 1; i > colStart; i-- {
				res = append(res, matrix[i][colStart])
			}
		}
		rowStart++
		rowEnd--

		colStart++
		colEnd--
	}
	return res
}

func main() {
	matrix := [][]int{
		//{1, 2, 3},
		//{4, 5, 6},
		//{7, 8, 9},
		//{10, 11, 12},
		//{1, 2, 3, 4},
		//{5, 6, 7, 8},
		//{9, 10, 11, 12},
		{1, 11},
		{2, 12},
		{3, 13},
		{4, 14},
		{5, 15},
		{6, 16},
		{7, 17},
		{8, 18},
		{9, 19},
		{10, 20},
	}

	//[[1,11],[2,12],[3,13],[4,14],[5,15],[6,16],[7,17],[8,18],[9,19],[10,20]]

	res := spiralOrder(matrix)
	fmt.Println(res)
}
