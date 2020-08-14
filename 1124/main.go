package main

import "fmt"

func main() {
	tests := [7][]int{
		{9, 9, 6, 0, 6, 6, 9},
		{8, 8, 0, 0, 7},
		{8, 8, 8, 8, 9},
		{9, 0, 0, 8, 7},
		{8, 9, 10, 7, 2},
		{8, 8, 8, 8, 9, 9},
		{9, 9, 6, 9, 8, 8, 0, 11, 8, 9, 9, 9},
	}
	for _, test := range tests {
		fmt.Println(longestWPI(test))
	}
}
func longestWPI(hours []int) int {
	var maxPeriodLen int
	var maxCurrentLen int

	var cycle bool
	var startPeriod, endPeriod int

	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			if cycle {
				maxCurrentLen++
			} else {
				cycle = true
				maxCurrentLen = 1
				startPeriod = i
			}
		} else {
			if cycle {
				maxCurrentLen--
				if maxCurrentLen == 0 {
					endPeriod = i
					cycle = false
					if endPeriod-startPeriod > maxPeriodLen {
						maxPeriodLen = endPeriod - startPeriod
					}
				}
			}
		}
	}
	if cycle {
		if maxCurrentLen > maxPeriodLen {
			maxPeriodLen = maxCurrentLen
		}
	}
	return maxPeriodLen
}
