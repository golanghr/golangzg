package main

// https://play.golang.org/p/vAPI9LyBmxP

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)

	const max = 5

	wg := sync.WaitGroup{}
	wg.Add(max)

	in := generator(max)

	for i := 0; i < max; i++ {
		go worker(&wg, in, i)
	}

	wg.Wait()
}

func worker(wg *sync.WaitGroup, in chan int, id int) {
	fmt.Println("start:", id)
	for job := range in {
		err := do(job)
		if err != nil {
			log.Printf("worker %v: %v", id, err)
			// log.Print(err)
			continue
		}
	}
	fmt.Println("stop:", id)
	wg.Done()
}

func do(n int) error {
	return fmt.Errorf("do fails for %v", n)
}

func generator(max int) chan int {
	in := make(chan int, max)
	go func() {
		for i := 0; i < max; i++ {
			in <- 100 + i
		}
		close(in)
	}()
	return in
}
