package main

// This program is sequential, race free but has deadlock.

// Deadlocks:
//   go run main.go
// Hangs:
//   go run -race main.go

func main() {
	var v int
	ch := make(chan int)
	v = 100
	select {
	case v = <-ch:
		// This branch writes to v
	case ch <- v:
		// This branch reads from v
	}
}
