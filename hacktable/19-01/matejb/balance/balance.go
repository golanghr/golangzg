package main

import (
	"fmt"
)

func worker(ch chan int) {
	for {
		<-ch
	}
}

func main() {
	chA := make(chan int)
	go worker(chA)

	chB := make(chan int)
	go worker(chB)

	chC := make(chan int)
	go worker(chC)

	type count struct {
		A int
		B int
		C int
	}

	var countNoBalance count
	for i := 0; i < 100; i++ {
		select {
		case chA <- i:
			countNoBalance.A++
		case chB <- i:
			countNoBalance.B++
		case chC <- i:
			countNoBalance.C++
		}
	}
	fmt.Printf("not balanced: %+v\n", countNoBalance)

	var countBalance count
	backup := [...]chan int{chA, chB, chC}
	for i := 0; i < 100; i++ {
		select {
		case chA <- i:
			countBalance.A++
			chA = nil
			chB = backup[1]
			chC = backup[2]
		case chB <- i:
			countBalance.B++
			chA = backup[0]
			chB = nil
			chC = backup[2]
		case chC <- i:
			countBalance.C++
			chA = backup[0]
			chB = backup[1]
			chC = nil
		}
	}
	fmt.Printf("balanced: %+v\n", countBalance)
}
