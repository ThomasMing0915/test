package main

import "fmt"

type point struct {
	posX int
	posY int
}

func movingCount(m int, n int, k int) int {
	mp := make(map[point]struct{}, 500)

	move(m, n, point{0, 0}, k, mp)
	//fmt.Println(mp)
	return len(mp)
}

//m：rows  n:columns   posX:row  posY:column
func move(m, n int, p point, k int, mp map[point]struct{}) {
	if p.posX < 0 || p.posX >= m {
		return
	}
	if p.posY < 0 || p.posY >= n {
		return
	}
	if digitalSum(p.posX)+digitalSum(p.posY) > k {
		return
	}
	if _, ok := mp[p]; ok {
		return
	}
	mp[p] = struct{}{}
	move(m, n, point{p.posX - 1, p.posY}, k, mp) //向上
	move(m, n, point{p.posX, p.posY + 1}, k, mp) //向右
	move(m, n, point{p.posX + 1, p.posY}, k, mp) //向下
	move(m, n, point{p.posX, p.posY - 1}, k, mp) //向左
}

func digitalSum(num int) int {
	sum := 0
	for num > 0 {
		sum = sum + (num % 10)
		num /= 10
	}

	return sum
}

func main() {
	for _, v := range [][3]int{{2, 3, 1}, {3, 3, 2}, {38, 15, 9}} {
		fmt.Println(movingCount(v[0], v[1], v[2]))
	}

}
