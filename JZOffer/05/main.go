package main

func replaceSpace(s string) string {
	rs := []rune(s)

	spaceCount := 0
	for i := 0; i < len(rs); i++ {
		if rs[i] == ' ' {
			spaceCount++
		}
	}

	news := make([]rune, 0, len(rs)+2*spaceCount)
	for i := 0; i < len(rs); i++ {
		if rs[i] == ' ' {
			news = append(news, '%', '2', '0')
			continue
		}
		news = append(news, rs[i])
	}

	return string(news)
}
