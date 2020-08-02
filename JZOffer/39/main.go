package main

func majorityElement(nums []int) int {
	m:=make(map[int]int,len(nums)/2)
	for _,num:=range nums{
		m[num]+=1
	}

	var maxCount,num int
	for k,v:=range m{
		if maxCount<v{
			maxCount=v
			num=k
		}
	}
	return num
}
