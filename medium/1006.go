package main

//clumsy * / + -

func clumsy(N int) int {
	stack := []int{N}
	ope := 1
	for i := N - 1; i > 0; i-- {
		switch ope % 4 {
		case 1:
			stack[len(stack)-1] *= i
		case 2:
			stack[len(stack)-1] /= i
		case 3:
			stack = append(stack, i)
		case 0:
			stack = append(stack, 0-i)
		}
		ope++
	}
	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}
