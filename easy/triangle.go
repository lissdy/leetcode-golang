package main

func maxProduct(nums []int) int {
	imax := make([]int, len(nums))
	imin := make([]int, len(nums))
	imax[0] = nums[0]
	imin[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		imax[i] = max(imax[i-1]*nums[i], max(nums[i], imin[i-1]*nums[i]))
		imin[i] = min(imin[i-1]*nums[i], min(nums[i], imax[i-1]*nums[i]))
		if imax[i] > ans {
			ans = imax[i]
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
