package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	if !IsNumber(s) {
		if strings.TrimSpace(s) == "" || !IsNumber(strings.TrimSpace(s)) {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	runes := []rune(s)
	var b, circle bool
	for i := 0; i < len(runes); i++ {
		if i == 0 {
			b = firstString(runes[i])
		} else if i == len(runes)-1 {
			b = lastString(runes[i])
		} else {
			b, circle = midString(runes[i-1], runes[i], circle)
		}
		if !b {
			return false
		}
	}
	return true
}

func firstString(r rune) bool {
	if r == '+' || r == '-' {
		return true
	}

	if r >= '0' && r <= '9' {
		return true
	}

	return false
}

func midString(pre, cur rune, circle bool) (bool, bool) {
	if circle && (cur == 'e' || cur == '.') {
		return false, true
	}
	if cur >= '0' && cur <= '9' {
		return true, circle
	}

	if cur == 'e' && (pre >= '0' && pre <= '9') {
		return true, true
	}

	if cur == 'E' && (pre >= '0' && pre <= '9') {
		return true, true
	}

	if cur == '.' && (pre >= '0' && pre <= '9') {
		return true, true
	}

	return false, circle
}

func lastString(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}

func main() {
	strs1 := []string{"+100", "5e2", "-123", "3.1416", "0123"}
	strs2 := []string{"12e", "1a3.14", "1.2.3", "+-5", "-1E-16", "12e+5.4"}

	strs3 := []string{" 12", "12 ", "11E1", "  "}

	for _, str := range strs1 {
		fmt.Println(isNumber(str))
	}

	for _, str := range strs2 {
		fmt.Println(isNumber(str))
	}

	for _, str := range strs3 {
		fmt.Println(isNumber(str))
	}
}
