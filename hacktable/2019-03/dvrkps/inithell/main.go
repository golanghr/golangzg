package main

import "fmt"

func init() {
	println("main/main: init before")
}

func main() {
	fmt.Printf("main/main: aa:%v az:%v za:%v zz:%v\n", aa, az, za, zz)
}

func init() {
	println("main/main: init after")
}
