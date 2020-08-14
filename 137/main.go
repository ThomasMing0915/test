package main

import "fmt"

func singleNumber(nums []int) int {
	//use hash
	numsmap := make(map[int]int, len(nums))
	for _, num := range nums {
		count, ok := numsmap[num]
		if ok {
			numsmap[num] = count + 1
		} else {
			numsmap[num] = 1
		}
	}

	for num, count := range numsmap {
		if count == 1 {
			return num
		}
	}

	//do not walk here
	return 0
}

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}
