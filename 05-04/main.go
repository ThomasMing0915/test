package main

import "fmt"

func main() {
	var val = 0x7fffffff + 2
	fmt.Println(val)
}

//func findClosedNumbers(num int) []int {
//	return nil
//}
//
//func swapBit(x, i, j int) int {
//	if ((x >> i) & 1) != ((x >> j) & 1) {
//		x ^= (1 << i) | (1 << j)
//	}
//	return x
//}
