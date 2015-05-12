package main

import (
	"encoding/json"
	"fmt"
)

type Gopher interface {
	Nod()
	Wink()
}

func play(gopher Gopher) {
	gopher.Nod()
	gopher.Wink()
}

// START 1 OMIT
type FakeGopher struct {
	Nodded bool
	Winked bool
}

func (g *FakeGopher) Nod() {
	g.Nodded = true // HL
}

func (g *FakeGopher) Wink() {
	g.Winked = true // HL
}

// END 1 OMIT

func (g *FakeGopher) ShowState() {
	jsonB, _ := json.MarshalIndent(g, "", "\t")
	fmt.Println(string(jsonB))
}

// START PLAY OMIT

func TestPlay() {
	g := &FakeGopher{}
	g.ShowState()
	play(g)
	g.ShowState()
}

// END PLAY OMIT
func main() {
	TestPlay()
}
