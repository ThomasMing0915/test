package main

func firstUniqChar(s string) byte {
	srune := []rune(s)
	mp := make(map[rune]int)

	for _, r := range srune {
		mp[r] = mp[r] + 1
	}

	for _, r := range srune {
		count, _ := mp[r]
		if count == 1 {
			return byte(r)
		}
	}
	return byte(' ')
}
