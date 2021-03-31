package main

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104
//
// Related Topics 数组 二分查找
// 👍 384 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	arr := []int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			arr = append(arr, matrix[i][j])
		}
	}
	return binSearch(arr, target)
}

func binSearch(arr []int, target int) bool {
	if arr[len(arr)/2] == target {
		return true
	}
	if len(arr) == 1 {
		return false
	}
	if arr[len(arr)/2] > target {
		binSearch(arr[:len(arr)/2], target)
	}
	if arr[len(arr)/2] < target {
		binSearch(arr[len(arr)/2:], target)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
