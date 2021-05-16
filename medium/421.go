package main

func findMaximumXOR(nums []int) int {
	ans := 0
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			tmp := nums[i] ^ nums[j]
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}
