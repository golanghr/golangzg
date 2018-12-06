package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := []int{1, 2, 3, 4, 5}

	ch := make(chan int, len(jobs))

	var wg sync.WaitGroup

	for _, v := range jobs {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ch <- v
		}(v)
	}

	wg.Wait()
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
