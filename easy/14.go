package main

func longestCommonPrefix(strs []string) string {
	ans := ""
	if len(strs) == 0 {
		return ans
	}
	for i := 0; i < len(strs[0]); i++ {
		for x := 1; x < len(strs); x++ {
			if len(strs[x]) <= i || strs[x][i] != strs[0][i] {
				return ans
			}
		}
		ans += string(strs[0][i])
	}
	return ans
}
