package main

func fairCandySwap(A []int, B []int) []int {
	// x,y   y = x + (a_sum+b_sum)/2
	a_sum, b_sum := 0, 0
	a_set := map[int]struct{}{}
	for _, v := range A {
		a_sum += v
		a_set[v] = struct{}{}
	}
	for _, v := range B {
		b_sum += v
	}
	tmp := (b_sum - a_sum) / 2
	for i := 0; ; i++ { // 答案肯定存在，不会死循环
		x := B[i] - tmp
		if _, has := a_set[x]; has {
			return []int{x, B[i]}
		}
	}
}
