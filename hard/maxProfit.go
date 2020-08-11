package main

import "fmt"

func main() {
	execute(sayHello,"Lisa")
}

func execute(hi func(value string),name string){
	hi(name)
}

func sayHello(name string){
	fmt.Println("Hello, I am " + name)
}