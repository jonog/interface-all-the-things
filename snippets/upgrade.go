package main

import "fmt"

type Gopher interface {
	Nod()
	Wink()
}

// START 1 OMIT
type Coder interface {
	Code()
}

// END 1 OMIT

func play(gopher Gopher) {
	gopher.Nod()
	gopher.Wink()
}

// START 2 OMIT
func writeCode(gopher Gopher) {

	// if our Gopher satisfies Coder, then type assert // HL
	if coder, ok := gopher.(Coder); ok {
		coder.Code()
	}
}

// END 2 OMIT
type SingGopher struct{}

func (g *SingGopher) Nod() {
	fmt.Println("can nod")
}

func (g *SingGopher) Wink() {
	fmt.Println("wink also can")
}

func (g *SingGopher) Code() {
	fmt.Println("can code")
}

// START PLAY OMIT
func main() {
	writeCode(&SingGopher{})
}

// END PLAY OMIT
