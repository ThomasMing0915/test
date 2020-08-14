package main

import (
	"fmt"
	"time"
)

func handle(i interface{}){
	fmt.Println(i)
}

func main(){
	tw:=New(time.Duration(1*time.Second),60,handle)
	tw.Start()

	tw.AddTimer(time.Duration(5*time.Second),"delay5","hello delay 5s")
	tw.AddTimer(time.Duration(10*time.Second),"delay10","hello delay 10s")
	tw.AddTimer(time.Duration(10*time.Second),"delay10-2","hello delay 10s")
	tw.AddTimer(time.Duration(61*time.Second),"delay61","hello delay 61s")

	time.Sleep(time.Duration(90*time.Second))
	tw.Stop()
}
