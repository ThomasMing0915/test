package main

func findIndex(main string,pattern string, hashValueFun func(string)int)int{
	if len(main)<len(pattern){
		return -1
	}

	patternHashValue:=0 //匹配串的hash值
	for i:=0;i<len(pattern);i++{
		patternHashValue+=hashValueFun(pattern[i:i+1])
	}

	mainHashValue:=0
	for i:=0;i<len(pattern);i++{
		mainHashValue+=hashValueFun(main[i:i+1])
	}

	if mainHashValue==patternHashValue{
		return 0
	}

	for i:=1;i<len(main);i++{
		//依次计算从 main[1:len(1+pattern)]的值
		if i+len(pattern)>len(main){
			break
		}
		mainHashValue=mainHashValue-hashValueFun(main[i-1:i])+hashValueFun(main[i+len(pattern)-1:i+len(pattern)])
		if mainHashValue==patternHashValue{
			return i
		}
	}
	return -1
}
