package main

import "fmt"

//暴力
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows := len(matrix)
	colums := len(matrix[0])

	ri := 0
	cj := colums - 1

	for ri < rows && cj >= 0 {
		if matrix[ri][cj] == target {
			return true
		} else if matrix[ri][cj] < target {
			ri++
		} else {
			cj--
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray2(matrix, 15))
	fmt.Println(findNumberIn2DArray2(matrix, 18))
	fmt.Println(findNumberIn2DArray2(matrix, 1))
	fmt.Println(findNumberIn2DArray2(matrix, 30))
	fmt.Println(findNumberIn2DArray2(matrix, 31))
}
