package main

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}

	tmpa, tmpb := 1, 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = (tmpa + tmpb) % 1000000007
		tmpa = tmpb
		tmpb = sum
	}
	return sum
}
