package main

func hammingDistance(x int, y int) int {
	tmp := x ^ y
	count := 0
	for tmp != 0 {
		if tmp&1 == 1 {
			count++
		}
		tmp >>= 1
	}
	return count
}
