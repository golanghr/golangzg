package main

// https://play.golang.org/p/M1BPVuZRoAJ

import (
	"fmt"
)

type Generic struct {
	Name string
}

func (c *Generic) Init() {
}

func (c *Generic) Clear() {
	c.Init()
}

func (c *Generic) GetName() string {
	return c.Name
}

type Cats struct {
	Generic
	Data []string
}

func (c *Cats) Init() {
	c.Data = []string{}
}

func main() {
	g := &Generic{Name: "Generic"}
	g.Init()
	fmt.Println(g.GetName())

	c := &Cats{}
	//c := &Cats{Name: "Cat"} // This does not work
	c.Name = "Cat"

	c.Init()
	c.Data = []string{"Buco"}
	fmt.Println(c.GetName(), c.Data)

	c.Clear()
	c.Data = append(c.Data, "Jerry")
	fmt.Println(c.GetName(), c.Data)
}
