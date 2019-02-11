package main

import "fmt"

func main() {
	fmt.Println(TheThing("WUT"))
}

func TheThing(name string) string {
	return "the thing " + name
}
