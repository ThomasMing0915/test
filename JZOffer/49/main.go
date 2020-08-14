package main

import "fmt"

func nthUglyNumber(n int) int {
	start := 1
	var res int
	for n > 0 {
		if isUglyNumber(start) {
			res = start
			n--
		}
		start++
	}
	return res
}

func isUglyNumber(n int) bool {
	if n == 1 {
		return true
	}

	for {
		flag := false
		if n%2 == 0 {
			n = n / 2
			flag = true
		}
		if n%3 == 0 {
			n = n / 3
			flag = true
		}
		if n%5 == 0 {
			n = n / 5
			flag = true
		}
		if !flag {
			return n == 1
		}
	}
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
