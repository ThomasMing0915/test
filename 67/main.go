package main

import "fmt"

func addBinary(a string, b string) string {
	arune := []rune(a)
	brune := []rune(b)

	maxLenAB := max(len(arune), len(brune)) + 1

	sumrune := make([]rune, maxLenAB)

	i, j, pos := 0, 0, maxLenAB
	flag := 0 //进位数
	for i < len(arune) && j < len(brune) {
		pos--
		sum := int((arune[len(arune)-1-i] - '0')) + int((brune[len(brune)-1-j] - '0')) + flag
		if sum%2 == 0 {
			sumrune[pos] = '0'
		} else {
			sumrune[pos] = '1'
		}
		flag = sum / 2

		i++
		j++
	}

	for i < len(arune) {
		pos--
		sum := int((arune[len(arune)-1-i] - '0')) + flag
		if sum%2 == 0 {
			sumrune[pos] = '0'
		} else {
			sumrune[pos] = '1'
		}
		flag = sum / 2
		i++
	}

	for j < len(brune) {
		pos--
		sum := int((brune[len(brune)-1-j] - '0')) + flag
		if sum%2 == 0 {
			sumrune[pos] = '0'
		} else {
			sumrune[pos] = '1'
		}
		flag = sum / 2
		j++
	}
	if flag > 0 {
		pos--
		sumrune[pos] = '1'
	}
	return string(sumrune[pos:])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	a := "0"
	b := "0"
	fmt.Println(addBinary(a, b))
}
