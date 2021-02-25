package main

//给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
//
// 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
//
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[1,4,7],[2,5,8],[3,6,9]]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3],[4,5,6]]
//输出：[[1,4],[2,5],[3,6]]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 1000
// 1 <= m * n <= 105
// -109 <= matrix[i][j] <= 109
//
// Related Topics 数组
// 👍 156 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func transpose(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	var ans = make([][]int, col)
	for i := 0; i < col; i++ {
		ans[i] = make([]int, row)
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			ans[i][j] = matrix[j][i]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
