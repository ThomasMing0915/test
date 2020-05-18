package main

func numTrees(n int) int {
	return catalan(n)
}

//卡特兰数
func catalan(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 5
	default:
		sum := 0
		for i := 0; i < n; i++ {
			sum += catalan(i) * catalan(n-i-1)
		}
		return sum
	}
}

func main() {
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Println(catalan(v))
	}
}
