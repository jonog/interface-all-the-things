package main

import "fmt"

type Gopher interface {
	Nod()
	Wink()
}

func play(gopher Gopher) {
	fmt.Println("---")
	gopher.Nod()
	gopher.Wink()
}

// START 1 OMIT
func pipeline(gophers ...Gopher) {
	for _, gopher := range gophers {
		play(gopher)
	}
}

// END 1 OMIT

type SingGopher struct{}

func (g *SingGopher) Nod() { // HL
	fmt.Println("can nod")
}

func (g *SingGopher) Wink() { // HL
	fmt.Println("wink also can")
}

type YodaGopher struct{}

func (g *YodaGopher) Nod() {
	fmt.Println("nod i can")
}

func (g *YodaGopher) Wink() {
	fmt.Println("wink i must")
}

// START PLAY OMIT
func main() {
	pipeline(&SingGopher{}, &YodaGopher{})
}

// END PLAY OMIT
