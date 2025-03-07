package main

import "fmt"

type Set[T comparable] = map[T]struct{}

func Add[T comparable](set Set[T], value T) Set[T] {
    set[value] = struct{}{}
    return set
}

func main() {
    set := Set[int]{
        0: {},
        1: {},
    }

    set = Add(set, 1)
    set = Add(set, 2)

    fmt.Printf("set is %T\n", set)
    fmt.Printf("set = %v\n", Values(set))
}

func Values[T comparable](set Set[T]) []T {
	var values []T
	for value := range set {
		values = append(values, value)
	}
	return values
}








