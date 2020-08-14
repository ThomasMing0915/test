package main

//import (
//	"fmt"
//	"log"
//	"runtime"
//	"time"
//)
//
//func readMemStats(){
//	var ms runtime.MemStats
//	runtime.ReadMemStats(&ms)
//
//	fmt.Printf("=====> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)\n",ms.Alloc,ms.HeapAlloc,ms.HeapReleased)
//}
//
//func test(){
//	container:=make([]int,0)
//
//	log.Println("====> loop begin.")
//	for i:=0;i<32*1000*1000;i++{
//		container=append(container,i)
//		if (i==16*1000*1000){
//			readMemStats()
//		}
//	}
//
//	log.Println("=====> loop end.")
//}
//
//func main(){
//	log.Println("====> [Start].")
//
//	readMemStats()
//	test()
//	readMemStats()
//
//	log.Println("=====>[force gc].")
//	runtime.GC()
//
//	go func(){
//		for{
//			readMemStats()
//			time.Sleep(10*time.Second)
//		}
//	}()
//
//	time.Sleep(3600*time.Second)
//}
