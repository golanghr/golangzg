package main

// https://play.golang.org/p/WdQCVp7rIU-

import "time"

func main() {
	const (
		capacity = 4
		max      = 8
	)
	ch := make(chan int, capacity)
	quit := make(chan struct{})

	go worker(ch, quit)
	for i := 0; i < max; i++ {
		select {
		case ch <- i:
			println("send", i)
		default:
			println("*** block", i)
			ch <- i
		}
	}
	close(ch)
	println("sending done")
	<-quit
}

func worker(ch chan int, quit chan struct{}) {
	for v := range ch {
		time.Sleep(1e9)
		println("receive", v)
	}
	close(quit)
}
