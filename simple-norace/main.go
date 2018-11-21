package main

// This program is NOT racy - parameters are passed by value

// go run -race main.go
// 0

import "fmt"

func main() {
	var v int
	ch := make(chan int)
	go s(ch, v)
	v++
	v = <-ch
	fmt.Println(v)
}

func s(ch chan int, v int) {
	ch <- v
}
