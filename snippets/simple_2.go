package main

import "fmt"

// START 1 OMIT
type Gopher interface {
	Nod()
	Wink()
}

func play(gopher Gopher) {
	fmt.Println("---")
	gopher.Nod()
	gopher.Wink()
}

// END 1 OMIT

type SingGopher struct{}

func (g *SingGopher) Nod() {
	fmt.Println("can nod")
}

// START PLAY OMIT
func main() {
	play(&SingGopher{})
}

// END PLAY OMIT
