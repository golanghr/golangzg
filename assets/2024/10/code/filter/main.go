package main

import "fmt"

func main() {
	for k, v := range shopping.Groceries {
		fmt.Println(k, v)
	}
	fmt.Println("====================")

	for i, g := range shopping.SortedByName() {
		fmt.Println(i, g)
	}
	fmt.Println("====================")

	for g := range shopping.Filter("fruit") {
		fmt.Println(g)
	}
}
