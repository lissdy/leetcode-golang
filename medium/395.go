package main

func longestSubstring(s string, k int) int {
	n := len(s)
	ans := 0
	for left := 0; left < n; left++ {
		dir := [26]int{}
		exist := []int{}
		for right := left; right < n; right++ {
			index := int(s[right] - 'a')
			dir[index] += 1
			exist = append(exist, index)
			flag := true
			for _, v := range exist {
				if dir[v] < k {
					flag = false
					break
				}
			}
			if flag {
				ans = max(ans, right-left+1)
			}
		}
	}
	return ans
}
