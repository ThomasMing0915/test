package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	prefixSum := make([]int, 0, target/2+1+1)
	prefixSum = append(prefixSum, 0)
	sum := 0
	for i := 1; i <= target/2+1; i++ {
		sum += i
		prefixSum = append(prefixSum, sum)
	}

	res := make([][]int, 0, 2)
	for i := 0; i < len(prefixSum); i++ {
		for j := i + 2; j < len(prefixSum); j++ {
			if prefixSum[j]-prefixSum[i] == target {
				res = append(res, slice(i, j))
			} else if prefixSum[j]-prefixSum[i] > target {
				break
			}
		}
	}
	return res
}

func slice(i, j int) []int {
	if j <= i {
		return nil
	}
	s := make([]int, 0, j-i)

	for index := i + 1; index <= j; index++ {
		s = append(s, index)
	}

	return s
}

func main() {
	for _, s := range []int{9, 15} {
		fmt.Println(findContinuousSequence(s))
	}
}
