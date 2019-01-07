// +build ignore

// This program is not racy thanks to the mutex

package main

import (
	"fmt"
	"sync"
)

func main() {
	var v int
	v = get() + 1
	mu := new(sync.Mutex)
	go s(mu, &v)
	mu.Lock()
	v++
	mu.Unlock()
	mu.Lock()
	fmt.Println(v)
	mu.Unlock()
}

func s(mu *sync.Mutex, ptr *int) {
	mu.Lock()
	*ptr = 10
	mu.Unlock()
}

func get() int { return 0 }
