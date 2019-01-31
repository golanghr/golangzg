package main

// https://play.golang.org/p/hYHMDbDcZFh

import (
	"log"
)

type Dog struct {
	Name      string
	BirthYear int
}

type Rabbit struct {
	_         [0]int
	Name      string
	BirthYear int
}

func main() {
	pluto := Dog{"Pluto", 1929}
	log.Println(pluto)

	//bugsBunny := Rabbit{"Bugs Bunny", 1938}
	bugsBunny := Rabbit{
		Name:      "Bugs Bunny",
		BirthYear: 1938,
	}
	log.Println(bugsBunny)
}
