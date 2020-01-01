package main

import "fmt"

func main() {
	var nums = []int{1,3,5,6}
	fmt.Println(searchInsert(nums,0))
}

func searchInsert(nums []int, target int) int {
	for i,_ := range nums {
		if i == len(nums)-1 && nums[i]<target {
			return len(nums)
		}
		if nums[i]<target && nums[i+1]<target {
			continue
		}
		if nums[i]==target {
			return i
		}
		if nums[i]<target && nums[i+1]>=target{
			return i+1
		}
	}
	return 0
}
