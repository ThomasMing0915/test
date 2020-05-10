package main

import "fmt"

func main() {
	tests := [5][2]int{
		{29, 15},
		{1, 1},
		{2, 1},
		{-2, -1},
		{-2, 1},
	}
	for _, test := range tests {
		fmt.Println(convertInteger(test[0], test[1]))
	}
}

func convertInteger(A int, B int) int {
	//异或运算得到有多个位置不一样，然后求1的个数
	xorAB := A ^ B

	count := 0 //统计1的数量
	for xorAB != 0 {
		if xorAB&0x80000000 == 0x80000000 {
			count++
		}
		xorAB <<= 1 //为什么不右移， 负数为右移高位会带符号位
	}
	return count
}
