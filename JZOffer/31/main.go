package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	tmp := make([]int, 0, len(pushed))
	var i, j int
	for i < len(pushed) && j < len(popped) {
		if len(tmp) > 0 && tmp[len(tmp)-1] == popped[j] {
			tmp = tmp[:len(tmp)-1]
			j++
		} else if pushed[i] != popped[j] {
			tmp = append(tmp, pushed[i])
			i++
		} else {
			i++
			j++
		}
	}

	for ; j < len(popped); j++ {
		if len(tmp) > 0 && tmp[len(tmp)-1] == popped[j] {
			tmp = tmp[:len(tmp)-1]
		} else {
			break
		}
	}

	return len(tmp) == 0
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3}, []int{3, 2, 1}))
}
