package main

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums)-1 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		i += 2
	}
	return nums[len(nums)-1]
}
