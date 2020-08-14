package main

import "fmt"

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := (left + right) / 2
		if mid-1 < 0 || mid+1 >= len(numbers) {
			return numbers[0]
		} else {
			if numbers[mid-1] > numbers[mid] && numbers[mid] < numbers[mid+1] {
				return numbers[mid]
			} else if numbers[mid] < numbers[0] {
				right = mid - 1
			} else if numbers[mid] >= numbers[0] {
				left = mid + 1
			}
		}
	}
	return -1
}

func main() {
	tests := [][]int{
		{3, 4, 5, 1, 2},
		{2, 2, 2, 0, 1},
		{1, 2, 3, 4, 5},
		{0, 1, 2, 2, 2},
		{2, 2, 2, 0, 1},
		{1},
	}
	for _, test := range tests {
		fmt.Println(minArray(test))
	}
}
