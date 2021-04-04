package main

func numRabbits(answers []int) int {
	dir := map[int]int{}
	for _, v := range answers {
		dir[v]++
	}
	ans := 0
	for k, v := range dir {
		ans += ((k + v) / (k + 1)) * (k + 1)
	}
	return ans
}

// 向上取整
// x+y/y+1
