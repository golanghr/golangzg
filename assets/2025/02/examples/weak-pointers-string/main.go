package main

import (
	"fmt"
	"runtime"
	"weak"
)

func PrintPtr(p *string) {
	if p == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(*p)
}

func main() {
    data := "TVZ"
    p := &data

    // Make creates a weak pointer from a pointer to some value of type T.
    // func Make[T any](ptr *T) Pointer[T] {
    w := weak.Make(p)

    d := w.Value()
    PrintPtr(d)

    runtime.GC()

    d = w.Value()
    PrintPtr(d)
}
