package main

//给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。
//
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
//
//
//
// 示例 1：
//
//
//输入：[1, 2, 2, 3, 1]
//输出：2
//解释：
//输入数组的度是2，因为元素1和2的出现频数最大，均为2.
//连续子数组里面拥有相同度的有如下所示:
//[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
//最短连续子数组[2, 2]的长度为2，所以返回2.
//
//
// 示例 2：
//
//
//输入：[1,2,2,3,1,4,2]
//输出：6
//
//
//
//
// 提示：
//
//
// nums.length 在1到 50,000 区间范围内。
// nums[i] 是一个在 0 到 49,999 范围内的整数。
//
// Related Topics 数组
// 👍 224 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

type item struct{ count, left, right int }

func findShortestSubArray(nums []int) int {
	dir := map[int]item{}
	degree, ans := 0, 0
	for i, v := range nums {
		if e, has := dir[v]; has {
			e.count++
			degree = max(degree, e.count)
			e.right = i
		} else {
			dir[v] = item{1, i, i}
		}
	}
	for _, e := range dir {
		if e.count == degree {
			ans = min(ans, e.right-e.left+1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
