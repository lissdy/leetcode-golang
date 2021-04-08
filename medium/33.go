package main

func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] { // 如果左侧是有序集合
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右侧是有序集合
			if nums[r] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
