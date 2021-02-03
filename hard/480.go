//中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
//
// 例如：
//
//
// [2,3,4]，中位数是 3
// [2,3]，中位数是 (2 + 3) / 2 = 2.5
//
//
// 给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗
//口中元素的中位数，并输出由它们组成的数组。
//
//
//
// 示例：
//
// 给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。
//
// 窗口位置                      中位数
//---------------               -----
//[1  3  -1] -3  5  3  6  7       1
// 1 [3  -1  -3] 5  3  6  7      -1
// 1  3 [-1  -3  5] 3  6  7      -1
// 1  3  -1 [-3  5  3] 6  7       3
// 1  3  -1  -3 [5  3  6] 7       5
// 1  3  -1  -3  5 [3  6  7]      6
//
//
// 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
//
//
//
// 提示：
//
//
// 你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。
// 与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。
//
// Related Topics Sliding Window
// 👍 147 👎 0

// 一看hard模式就知道不是那么简单，果然第30测试用例超出时间限制
package main

//TODO 超时，用大小堆实现
import "sort"

func main() {
	nums := []int{1, 4, 2, 3}
	medianSlidingWindow(nums, 4)
}

func medianSlidingWindow(nums []int, k int) []float64 {
	ans := []float64{}
	sub_arr := [][]int{}
	for left := 0; left <= len(nums)-k; left++ {
		sub := []int{}
		for right := left; right < left+k; right++ {
			sub = append(sub, nums[right])
		}
		sub_arr = append(sub_arr, sub)
	}
	for _, v := range sub_arr {
		sort.Ints(v)
		var mid float64
		if k%2 == 0 {
			mid = float64((v[k/2-1] + v[k/2])) / 2
		} else {
			mid = float64(v[k/2])

		}
		ans = append(ans, mid)
	}
	return ans
}
