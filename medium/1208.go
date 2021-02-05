package main

func equalSubstring(s string, t string, maxCost int) int {
	res := 0
	diff := []int{}
	for i := range s {
		diff = append(diff, abs(int(s[i])-int(t[i])))
	}
	for left := range s {
		sum := 0
		for right := left; right < len(s); right++ {
			sum += diff[right]
			if sum > maxCost {
				break
			} else {
				res = max(res, right+1-left)
			}
		}

	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return 0 - x
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
