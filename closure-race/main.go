package main

// This program contains a race on variable x not allocated in heap
// and does not explicitly use pointer type in parameter.

import "fmt"

func main() {
	var x int
	ch := make(chan int)
	go func(ch chan int) {
		x = 10
		ch <- 1
	}(ch)
	x = 11
	<-ch
	fmt.Println(x)
}
