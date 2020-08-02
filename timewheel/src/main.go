package main

import (
	"fmt"
	"time"
)

func handle(i interface{}){
	fmt.Println(i)
}

func main(){
	tw:=New(1,60,handle)
	tw.Start()

	tw.AddTimer(5,"delay5","hello delay 5s")
	tw.AddTimer(10,"delay10","hello delay 10s")
	tw.AddTimer(10,"delay10-2","hello delay 10s")
	tw.AddTimer(61,"delay61","hello delay 61s")

	time.Sleep(time.Duration(90*time.Second))
	tw.Stop()
}
