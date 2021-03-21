package main

import "fmt"

func main() {
	var word string = "hello"
	for _, v := range word {
		fmt.Println(string(v))
	}
}

func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}

func climbStairs(n int) int {
	mem := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i <= 2 {
			mem[i] = i
		} else {
			mem[i] = mem[i-1] + mem[i-2]
		}
	}
	return mem[n]
}
