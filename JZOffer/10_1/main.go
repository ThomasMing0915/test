package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	tmpa, tmpb := 0, 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = (tmpa + tmpb) % 1000000007
		tmpa = tmpb
		tmpb = sum
	}
	return sum
}
