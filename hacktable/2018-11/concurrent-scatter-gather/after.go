package main

import "fmt"

func main() {
	jobs := []int{1, 2, 3, 4, 5}

	ch := make(chan int, len(jobs))

	for _, v := range jobs {
		go func(v int) {
			ch <- v
		}(v)
	}

	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}
