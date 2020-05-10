package main

import "fmt"

func hash(s string)int{
	bs:=[]byte(s)
	ba:=[]byte("a")
	if len(bs)==0 || len(ba)==0{
		return 0
	}
	return int(bs[0]-ba[0])
}

func main(){
	fmt.Println(findIndex("aaab","b",hash))
}
