package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	//from tail to head
	pos := m + n - 1
	i := m - 1
	j := n - 1

	for j >= 0 {
		//why add i<0   input [0] 0  [1] 1
		if i < 0 || nums2[j] >= nums1[i] {
			nums1[pos] = nums2[j]
			pos--
			j--
		} else {
			nums1[pos] = nums1[i]
			pos--
			i--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
