package main

//给定一个二进制数组， 计算其中最大连续1的个数。
//
// 示例 1:
//
//
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
//
//
// 注意：
//
//
// 输入的数组只包含 0 和1。
// 输入数组的长度是正整数，且不超过 10,000。
//
// Related Topics 数组
// 👍 194 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
	left, right, ans := 0, 0, 0
	n := len(nums)
	for right < n {
		if nums[right] == 1 {
			ans = max(ans, right-left+1)
			right++
		} else {
			if right < n-1 {
				left = right + 1
				right = left
			} else {
				right++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)