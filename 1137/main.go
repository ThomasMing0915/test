package main

func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	s0, s1, s2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		tmp := s2 + s1 + s0
		s0 = s1
		s1 = s2
		s2 = tmp
	}
	return s2
}
