package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	comparePos := -1

	///!!! note
	if len(strs) == 0 {
		return commonPrefix
	}

	for {
		comparePos += 1
		var char rune
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) <= comparePos {
				goto Exit
			}
			if i == 0 {
				rstr := []rune(strs[i])
				char = rstr[comparePos]
			}
			rstr := []rune(strs[i])
			if rstr[comparePos] != char {
				goto Exit
			}
		}
		commonPrefix = commonPrefix + string(char)
	}

Exit:
	return commonPrefix
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
