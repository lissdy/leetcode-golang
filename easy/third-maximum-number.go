//https://leetcode-cn.com/problems/third-maximum-number
package main

import "fmt"

func main() {
	var array = []int{1, 2, 2, 5, 3, 5}
	fmt.Println(thirdMax(array))
}

func thirdMax(nums []int) int {
	return nMax(nums, 3)
}

func nMax(s []int, n int) int {
	s = removeDul(s)
	if len(s) == 1 {
		return s[0]
	}
	mark := s[len(s)-1]
	var lowPart []int
	var midPart int
	var hihPart []int
	for _, x := range s {
		switch {
		case x < mark:
			lowPart = append(lowPart, x)
		case x == mark:
			midPart = x
		case x > mark:
			hihPart = append(hihPart, x)
		}
	}

	if len(hihPart) <= n-1 {
		return midPart
	}

	if len(hihPart) > n-1 {
		nMax(append(hihPart), n)
	} else {
		fmt.Println(lowPart)
		fmt.Println(n - len(hihPart) - 1)
		nMax(append(lowPart), n-len(hihPart)-1)
	}

	return s[0]
}

func removeDul(nums []int) []int {
	dic := map[int]int{}
	var res []int
	for _, x := range nums {
		if dic[x] == 1 {
			continue
		}
		dic[x] = 1
		res = append(res, x)
	}
	return res
}
