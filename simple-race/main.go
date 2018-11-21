package main

// This program is racy - despite output is consistently 1

import (
	"fmt"
)

func main() {
	var v int
	ch := make(chan int)
	go s(ch, &v)
	// time.Sleep(time.Millisecond) // uncomment to get v = 0
	v++ // Read/Write v
	v = <-ch
	fmt.Println(v)
}

func s(ch chan int, v *int) {
	ch <- *v // Read v
}
