package main

func characterReplacement(s string, k int) int {
	left, maxVaule := 0, 0
	dir := [26]int{}
	for right, v := range s {
		dir[v-'A']++
		maxVaule = max(maxVaule, dir[v-'A'])
		if right-left+1-maxVaule > k {
			dir[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
