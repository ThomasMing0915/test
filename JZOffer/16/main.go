package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	res := combine(x, n)
	if negative {
		return 1 / res
	}
	return res
}

func combine(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == 2 {
		return x * x
	}

	half := myPow(x, n/2)
	full := half * half
	if n%2 == 1 {
		full = full * x
	}
	return full
}

func main() {
	fmt.Println(myPow(2.0, -2))
}
