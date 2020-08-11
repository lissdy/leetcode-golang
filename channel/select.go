package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i ++
		}
	}()
	return out
}

func main(){
	var c1,c2 = generator(), generator()
	for{
		select {
		case n:= <- c1:
			fmt.Println("received from c1 :", n)
		case n:= <- c2:
			fmt.Println("c2:",n)
		default:
			time.Sleep(time.Second)
			fmt.Println("no receive")
		}
	}

}
