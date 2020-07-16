package main

func sumNums(n int) int {

	//这里还是使用了if  怎么才能不使用if
	if n == 1 {
		return n
	}
	return n + sumNums(n-1)
}
