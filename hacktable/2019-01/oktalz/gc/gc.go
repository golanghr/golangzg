package main

// https://play.golang.org/p/r0P4eFKmBqB

import (
	"log"
	"runtime"
	"time"
)

func example1() {
	go func() {
		for {
		}
	}()
	go func() {
		i := 0
		for {
			i++
			log.Println("random text", i)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	log.Println("Hi gophers")
	runtime.Gosched() //force running of goroutine
	log.Println("no GC call")
	// no GC, all ok, we exit
}

func example2() {
	//this example locks all threads
	go func() {
		for {
			//doing something intensive
		}
	}()
	go func() {
		i := 0
		for {
			i++
			log.Println("random text", i)
			time.Sleep(time.Millisecond * 10)
			//doing something that include for instance writing to disc
			runtime.Gosched() //just to make sure that this does not block
		}
	}()
	log.Println("Hi gophers")
	runtime.Gosched() //force running of goroutine
	log.Println("at one point GC kicks in")
	// and were are stuck here forever
	// both routines are stuck after some random text printing
	runtime.GC()
	// we never reach here
	log.Println("Bye all")
}

func example3() {
	go func() {
		for {
			//doing something intensive
			runtime.Gosched()
		}
	}()
	go func() {
		i := 0
		for {
			i++
			log.Println("random text", i)
			time.Sleep(time.Millisecond * 10)
			//runtime.Gosched() //we do not need that here
		}
	}()
	log.Println("Hi gophers")
	runtime.Gosched() //force running of goroutine
	log.Println("at one point GC kicks in")
	runtime.GC()
	log.Println("Bye all")
	//all ok :)
}

func main() {
	log.Println("\n----------\nWithout GC\n----------")
	example1() /* */

	/*log.Println("\n----------\nWith GC block\n----------")
	example2() /* */

	/*log.Println("\n----------\nWith GC but no block\n----------")
	example3() /* */

	log.Println("Done.")
}
