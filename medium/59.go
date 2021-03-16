package main

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	if n == 0 {
		return ans
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	for num != n*n+1 {
		//从左到右
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		//从上到下
		for i := top + 1; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		//从右到左
		for i := right - 1; i >= left; i-- {
			ans[bottom][i] = num
			num++
		}
		//从下到上
		for i := bottom - 1; i > top; i-- {
			ans[i][left] = num
			num++
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}
