package main

func findMin(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
