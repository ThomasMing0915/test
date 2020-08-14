package main

import "fmt"

//defer的执行顺序

func fa(){
	fmt.Println("function a")
}

func fb(){
	fmt.Println("function b")
}

func fc(){
	fmt.Println("function c")
}

//func main(){
//	defer fa()
//	defer fb()
//	defer fc()
//}
/*
function c
function b
function a
*/

//defer与return 谁先谁后
func fd()int{
	fmt.Println("function d")
	return 0
}

func fe() int{
	defer fc()
	return fd()
}

//func main(){
//	fe()
//}
/*
function d
function c

*/

func ff(i int)(r int){
	fmt.Println("r=",r)
	return 1
}

//func main(){
//	ff(2)
//}

//r= 0


func fg()(r int){
	defer func(){
		r=r*5
	}()

	return 1
}

//func main(){
//	fmt.Println(fg())
//}
//5

//defer 遇见panic

func defer_call(){
	defer func(){
		fmt.Println("defer:panic之前1")
	}()

	defer func(){
		fmt.Println("defer:panic之前2")
	}()

	panic("异常内容") //触发defer出栈

	defer func(){
		fmt.Println("defer:panic之后,永远执行不到")
	}()
}

//func main(){
//	defer_call()
//
//	fmt.Println("main 正常结束")
//}

/*
defer:panic之前2
defer:panic之前1
panic: 异常内容

goroutine 1 [running]:
main.defer_call()
        /Users/mingyong/WorkSpace/test/defer/main.go:86 +0x68
main.main()
        /Users/mingyong/WorkSpace/test/defer/main.go:94 +0x22


*/

//defer遇见panic 并捕获异常

func defer_call2(){
	defer func(){
		fmt.Println("defer:panic之前1,捕获异常")
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()

	defer func(){
		fmt.Println("defer:panic之前2,不捕获")
	}()

	panic("异常内容") //触发defer出栈

	defer func(){
		fmt.Println("defer:panic之后,永远执行不到")
	}()
}

//func main(){
//	defer_call2()
//
//	fmt.Println("main 正常结束")
//}

/*
defer:panic之前2,不捕获
defer:panic之前1,捕获异常
异常内容
main 正常结束
*/

//defer中包含panic


//func main(){
//	defer func(){
//		if err:=recover();err!=nil{
//			fmt.Println(err)
//		}else{
//			fmt.Println("fatal")
//		}
//	}()
//
//	defer func(){
//		panic("defer panic")
//	}()
//
//	panic("panic")
//}

/*
defer panic
 */

//defer下的函数参数包含子函数

func f(i ,v int)int{
	fmt.Println(i)
	return i
}

func main(){
	defer f(1,f(3,0))
	defer f(2,f(4,0))
}

/*
3
4
2
1

*/

