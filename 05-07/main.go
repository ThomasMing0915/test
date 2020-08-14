package main

func exchangeBits(num int) int {
	changeNum := num
	for i := 0; i < 32; i += 2 {
		changeNum = swapBit(changeNum, i, i+1)
	}
	return changeNum
}

func swapBit(x, i, j int) int {
	if ((x >> i) & 1) != ((x >> j) & 1) {
		x ^= (1 << i) | (1 << j)
	}
	return x
}
