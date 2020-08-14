package main

import "fmt"

func constructArr(a []int) []int {
	ca:=make([]int,0,len(a)+1)
	cb:=make([]int,0,len(a)+1)

	ca=append(ca,1)
	cb=append(cb,1)

	multi:=1
	for i:=0;i<len(a);i++{
		multi*=a[i]
		ca=append(ca,multi)
	}

	multi=1
	for i:=len(a)-1;i>=0;i--{
		multi*=a[i]
		cb=append(cb,multi)
	}

	res:=make([]int,0,len(a))
	res=append(res,cb[len(cb)-1])
	for i:=1;i<len(a);i++{
		res=append(res,ca[i-1]*cb[len(cb)-2])
	}

	return res
}

func main(){
	fmt.Println(1^1)
}
