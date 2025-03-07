package main

import (
	"fmt"
	"runtime"
	"weak"
)

type Location struct {
	Name                string
	Latitude, Longitude float32
}

func PrintLocation(p *Location) {
	if p == nil {
		fmt.Println("nil")
		return
	}
	fmt.Printf("%s: %f, %f\n", p.Name, p.Latitude, p.Longitude)
}

func main() {
    p := &Location{
        Name:      "TVZ",
        Latitude:  45.810955326640254,
        Longitude: 16.04167597609013,
    }
    w := weak.Make(p)

    d := w.Value()
    PrintLocation(d)
    _ = p // not really necessary

    runtime.GC()

    d = w.Value()
    PrintLocation(d)
}
