package main

func romanToInt(s string) int {
	dir := map[string]int{}
	dir["I"] = 1
	dir["V"] = 5
	dir["X"] = 10
	dir["L"] = 50
	dir["C"] = 100
	dir["D"] = 500
	dir["M"] = 1000
	dir["IV"] = 4
	dir["IX"] = 9
	dir["XL"] = 40
	dir["XC"] = 90
	dir["CD"] = 400
	dir["CM"] = 900
	ans := 0
	flag := false
	for i := 0; i < len(s)-1; i++ {
		if v, exist := dir[string(s[i])+string(s[i+1])]; exist {
			ans += v
			i++
			if i == len(s)-1 {
				flag = true
			}
		} else {
			ans += dir[string(s[i])]
		}
	}
	if !flag {
		ans += dir[string(s[len(s)-1])]
	}
	return ans
}
