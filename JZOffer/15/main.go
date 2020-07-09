package main

import "fmt"

func hammingWeight(num uint32) int {
	var v uint32
	v = 0x00000001
	count := 0
	for i := 0; i < 32; i++ {
		vv := v << i
		if num&vv != 0 {
			count++
		}
	}
	return count
}

func main() {
	for _, v := range []uint32{1023} {
		fmt.Println(hammingWeight(v))
	}
}
