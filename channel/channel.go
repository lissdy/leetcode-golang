package main

import (
	"fmt"
)

func worker(worker int) chan<- int {
	c := make(chan int)
	go func() {
		for n := range c {
			fmt.Printf("worker %d received %c \n", worker, n)
		}
	}()
	return c
}

func chanDemo() {
	for i := 0; i < 10; i++ {
		cc := worker(i)
		cc <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		cc := worker(i)
		cc <- 'A' + i
	}
}

func main() {
	chanDemo()
	//time.Sleep(time.Second)
}
