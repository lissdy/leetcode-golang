package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan int,2)
	go func() {
		for{
			n,ok:= <- c
			if !ok{
				break
			}
			fmt.Print(n)
		}
	}()
	c <- 1
	c <- 2
	close(c)
	time.Sleep(time.Second)
}
