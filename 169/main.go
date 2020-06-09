package main

func majorityElement(nums []int) int {
	numscountMap := make(map[int]int, 8)
	for i := 0; i < len(nums); i++ {
		count, ok := numscountMap[nums[i]]
		if !ok {
			numscountMap[nums[i]] = 1
			continue
		}
		numscountMap[nums[i]] = count + 1
	}
	maxCount := 0
	maxCountNum := nums[0]

	for k, v := range numscountMap {
		if v > maxCount {
			maxCount = v
			maxCountNum = k
		}
	}
	return maxCountNum
}
