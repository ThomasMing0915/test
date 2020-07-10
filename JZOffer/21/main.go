package main

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if !isEven(nums[left]) {
			left++
			continue
		}
		if isEven(nums[right]) {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	return nums
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
