package main

//给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
//
//
//上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
//
//
//
// 示例：
//
//
//给定 matrix = [
//  [3, 0, 1, 4, 2],
//  [5, 6, 3, 2, 1],
//  [1, 2, 0, 1, 5],
//  [4, 1, 0, 1, 7],
//  [1, 0, 3, 0, 5]
//]
//
//sumRegion(2, 1, 4, 3) -> 8
//sumRegion(1, 1, 2, 2) -> 11
//sumRegion(1, 2, 2, 4) -> 12
//
//
//
//
// 提示：
//
//
// 你可以假设矩阵不可变。
// 会多次调用 sumRegion 方法。
// 你可以假设 row1 ≤ row2 且 col1 ≤ col2 。
//
// Related Topics 动态规划
// 👍 200 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	preSum []int
	col    int
}

func Constructor(matrix [][]int) NumMatrix {
	n := 0
	if len(matrix) > 0 {
		n = len(matrix[0])
	}
	sum := make([]int, len(matrix)*n+1)
	for i, v := range matrix {
		for j, vv := range v {
			index := i*n + j
			sum[index+1] = sum[index] + vv
		}
	}
	return NumMatrix{sum, n}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for i := row1; i <= row2; i++ {
		begin := i*this.col + col1
		end := i*this.col + col2
		ans += this.preSum[end+1] - this.preSum[begin]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
