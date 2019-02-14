package main

// https://play.golang.org/p/n8oWtqHwrJY

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	all := []person{
		{name: "Art", legal: false, female: false},
		{name: "Betty", legal: false, female: true},
		{name: "Julia", legal: true, female: true},
		// {name: "", legal: true, female: true},
		{name: "Romeo", legal: true, female: false},
	}

	g := group{ps: all}

	result := g.valid().legals().females()
	if result.err != nil {
		log.Fatal(result.err)
	}

	for _, p := range result.ps {
		fmt.Println(p.name)

	}
}

type person struct {
	name   string
	legal  bool
	female bool
}

type group struct {
	ps  []person
	err error
}

func (g group) valid() group {
	if g.err != nil {
		return g
	}
	for _, p := range g.ps {
		if p.name == "" {
			g.err = errors.New("empty name")
			break
		}
	}
	return g
}

func (g group) legals() group {
	all := []person{}
	if g.err != nil {
		return g
	}
	for _, p := range g.ps {
		if p.legal {
			all = append(all, p)
		}
	}
	g.ps = all
	return g
}

func (g group) females() group {
	all := []person{}
	if g.err != nil {
		return g
	}
	for _, p := range g.ps {
		if p.female {
			all = append(all, p)
		}
	}
	g.ps = all
	return g
}
