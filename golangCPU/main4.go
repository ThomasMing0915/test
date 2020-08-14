package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"time"
	_ "net/http/pprof"
)

func test(){
	log.Println("====> loop begin.")

	for i:=0;i<1000;i++{
		log.Println(genSomeBytes())
	}

	log.Println("=====> loop end.")
}

//生产一个随机字符串
func genSomeBytes() *bytes.Buffer{
	var buff bytes.Buffer

	for i:=1;i<20000;i++{
		buff.Write([]byte{'0'+byte(rand.Intn(10))})
	}

	return &buff
}

func main(){
	go func(){
		for{
			test()
			time.Sleep(time.Second*1)
		}

	}()

	http.ListenAndServe("0.0.0.0:10001",nil)
}
