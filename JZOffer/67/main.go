package main

import "fmt"

func strToInt(str string) int{
	if str==""{
		return 0
	}

	r:=[]rune(str)
	var i int
	var breakFlag bool
	positive:=true

	for i=0;i<len(r);i++{
		switch r[i] {
		case '+':
			i++
			breakFlag=true
		case '-':
			i++
			positive=false
			breakFlag=true
		case ' ':
		case '0':
			//return 0 面试现场要与面试官确认  本题就不能直接return   00000123  要返回123
			breakFlag=true
		case '1':
			breakFlag=true
		case '2':
			breakFlag=true
		case '3':
			breakFlag=true
		case '4':
			breakFlag=true
		case '5':
			breakFlag=true
		case '6':
			breakFlag=true
		case '7':
			breakFlag=true
		case '8':
			breakFlag=true
		case '9':
			breakFlag=true
		default:
			return 0
		}

		if breakFlag{
			break
		}
	}

	res:=0
	for ;i<len(r);i++{
		if r[i]>='0' && r[i]<='9'{
			res=res*10+int(r[i]-'0')
			//为啥要在这里就要加判断 这里是debug补上的 开始并没有做对
			if !positive{
				if res>(0x7fffffff+1){
					break
				}
			}else{
				if res>0x7fffffff  || res<0{
					break
				}
			}
		}else{
			break
		}
	}

	if !positive{
		fmt.Println("yes")
		if res>(0x7fffffff+1){
			res=(0x7fffffff+1)
		}
		res=-1*res
	}else{
		if res>0x7fffffff  || res<0{
			res=0x7fffffff
		}
	}
	return res
}

func main(){
   //for _, s:=range []string{"42","   -42","4193 with words","words and 987"}{
   //		fmt.Println(strToInt(s))
   //}
   //fmt.Println(0x7fffffff)

   fmt.Println(strToInt("9223372036854775808"))
}
