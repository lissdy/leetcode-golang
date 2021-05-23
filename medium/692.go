package main

import "fmt"

func topKFrequent(words []string, k int) []string {
	count := map[string]int{}
	for _, v := range words {
		count[v]++
	}
	for item := range count {
		fmt.Println(item)
	}
	ans := []string{}
	return ans
}
