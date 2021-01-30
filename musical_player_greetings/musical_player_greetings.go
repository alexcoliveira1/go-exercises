package main

import (
	"fmt"
)

type trumpter struct {
	Name string
}
type violinist struct {
	Name string
}

func (t trumpter) Greeting() {
	fmt.Printf("My name is %v and I'm a trumpter\n", t.Name)
}

func (v violinist) Greeting() {
	fmt.Printf("My name is %v and I'm a violinist\n", v.Name)
}

type musicalPlayer interface {
	Greeting()
}

func main() {
	players := []musicalPlayer{
		trumpter{"Bob"},
		violinist{"Jon"},
	}
	for _, player := range players {
		player.Greeting()
	}
	fmt.Printf("MAIN\n")
}
