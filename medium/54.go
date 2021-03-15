package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	ans := []int{}
	row := len(matrix)
	col := len(matrix[0])
	left, right, top, bottom := 0, col-1, 0, row-1
	for left <= right && top <= bottom {
		// 从左到右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		// 从上到下
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if left < right && top < bottom {
			//从右到左
			for i := right - 1; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			//从下到上
			for i := bottom - 1; i >= top+1; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}
