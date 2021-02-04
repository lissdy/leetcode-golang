//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
//
//
//
// 示例：
//
//
//输入：[1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 30,000。
// 所给数据范围 [-10,000，10,000]。
//
// Related Topics 数组
// 👍 133 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	init_sum := 0
	for i := 0; i < k; i++ {
		init_sum += nums[i]
	}
	max_sum := init_sum
	for left := 1; left <= len(nums)-k; left++ {
		right := left + k - 1
		init_sum = init_sum - nums[left-1] + nums[right]
		max_sum = max(max_sum, init_sum)
	}
	return float64(max_sum) / float64(k)
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
