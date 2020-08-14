package main

var res []int

func getLeastNumbers(arr []int, k int) []int {
	res = make([]int, 0, k)

	recursive(arr, 0, len(arr)-1, k)

	return res
}

func recursive(arr []int, start, end, k int) {
	if start > end {
		return
	}

	index := partation(arr, start, end)
	if index == -1 {
		return
	}

	if index == (start + k) {
		res = append(res, arr[start:index]...)
		return
	} else if index < (start + k) {
		recursive(arr, index, end, start+k-index)
	} else {
		recursive(arr, start, index, k)
	}
}

func partation(arr []int, start, end int) int {
	if start >= len(arr) || end >= len(arr) {
		return -1
	}

	if end > start {
		return -1
	}

	pivot := arr[start]

	for start < end {
		for ; end > start; end-- {
			if arr[end] < pivot {
				arr[start] = arr[end]
				break
			}
		}

		for ; start < end; start++ {
			if arr[start] > pivot {
				arr[end] = arr[start]
				break
			}
		}
	}

	return start
}

//func getLeastNumbers(arr []int, k int) []int {
//	sort.Slice(arr,func(i,j int)bool{
//		return arr[i]<=arr[j]
//	})
//	return arr[:k]
//}
