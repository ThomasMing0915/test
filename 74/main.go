package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0])==0{
		return false
	}

	//锁定行
	startRows,endRows:=0,len(matrix)-1

	var midRows int

	for startRows<=endRows{
		midRows=(endRows-startRows)/2+startRows

		if matrix[midRows][0]<=target && matrix[midRows][len(matrix[midRows])-1]>=target{
			break
		}else if matrix[midRows][0]>target{
			endRows-=1
		}else {
			startRows+=1
		}
	}

	if startRows> endRows{
		return false
	}

	//行内二分查找
	left,right:=0,len(matrix[midRows])-1

	for left<=right{
		mid:=(right-left)/2+left

		if matrix[midRows][mid]==target{
			return true
		} else if matrix[midRows][mid]>target{
			right--
		}else{
			left++
		}
	}

	return false
}

func main(){
	maxtrix:=[][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,50},
	}

	fmt.Println(searchMatrix(maxtrix,13))


}
