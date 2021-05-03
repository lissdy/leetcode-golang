package main

import "strconv"

func countAndSay(n int) string {
	result := "1"
	for i := 2; i <= n; i++ {
		result = count(result)
	}
	return result
}

func count(s string) string {
	pre, count := string(s[0]), 1
	ans := ""
	for i := 1; i < len(s); i++ {
		if pre == string(s[i]) {
			count++
		} else {
			ans += strconv.Itoa(count) + pre
			pre = string(s[i])
			count = 1
		}
	}
	ans += strconv.Itoa(count) + pre
	return ans
}
