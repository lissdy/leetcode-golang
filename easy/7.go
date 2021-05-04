package main

import "math"

func reverse(x int) int {
	nevFlag := false
	if x < 0 {
		nevFlag = true
		x = 0 - x
	}
	rev := 0
	for x != 0 {
		tmp := x % 10
		x /= 10
		rev = rev*10 + tmp
	}
	if rev > math.MaxInt32 {
		rev = 0
	}
	if nevFlag {
		rev = 0 - rev
	}
	return rev
}
