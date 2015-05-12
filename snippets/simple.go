package main

import "fmt"

// START 1 OMIT
type Gopher interface {
	Nod()
	Wink()
}

// END 1 OMIT
// START 2 OMIT
func play(gopher Gopher) {
	gopher.Nod()
	gopher.Wink()
}

// END 2 OMIT
// START 3 OMIT
type SingGopher struct{}

func (g *SingGopher) Nod() {
	fmt.Println("can nod")
}

func (g *SingGopher) Wink() {
	fmt.Println("wink also can")
}

// END 3 OMIT

// START PLAY OMIT
func main() {
	play(&SingGopher{})
}

// END PLAY OMIT
