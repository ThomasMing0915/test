package main

func lemonadeChange(bills []int) bool {
	money := make([]int, 3) //money[0]--20  money[1]--10 money[2]==5
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			money[2]++
		} else if bills[i] == 10 {
			if money[2] < 1 {
				return false
			}
			money[1]++
			money[2]--
		} else if bills[i] == 20 {
			//note!!! why add if money[2]==0
			if money[2] == 0 || money[1]*10+money[2]*5 < 15 {
				return false
			}
			money[0]++
			if money[1] >= 1 {
				money[1]--
				money[2]--
			} else {
				money[2] -= 3
			}
		}
	}
	return true
}
