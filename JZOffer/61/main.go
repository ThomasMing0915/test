package main

import "fmt"

var cardMap = map[int]int{
	0:  0,
	65: 1,
	2:  2,
	3:  3,
	4:  4,
	5:  5,
	6:  6,
	7:  7,
	8:  8,
	9:  9,
	10: 10,
	74: 11,
	81: 12,
	75: 13,
}

func isStraight(nums []int) bool {
	array := make([]int, 15, 15)
	king := 0
	for _, num := range nums {
		//v, ok := cardMap[num]
		//if !ok {
		//	panic("not exist")
		//}
		v := num
		if num == 0 {
			king++
			continue
		}
		array[v] = 1
	}

	if king > 0 {
		start := 0
		for i := 0; i < len(array); i++ {
			if array[i] != 0 {
				start = i
				break
			}
		}

		for king > 0 {
			if array[start] == 0 {
				array[start] = 1
				king--
			}
			start++
		}

	}

	//startIndex, endIndex := -1, -1
	//for i := 0; i < len(array); i++ {
	//	if array[i] != 0 {
	//		if startIndex == -1 {
	//			startIndex = i
	//		} else {
	//			endIndex = i
	//		}
	//	}
	//}

	start := 0
	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			start = i
			break
		}
	}

	i := 0
	for i = start; i < start+5 && i < len(array); i++ {
		if array[i] == 0 {
			return false
		}
	}

	return i-start >= 5
}

func main() {
	fmt.Println(isStraight([]int{11, 8, 12, 8, 10}))
}
