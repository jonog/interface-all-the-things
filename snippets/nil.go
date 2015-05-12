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

type SingGopher struct{}

func (g *SingGopher) Nod() {
	fmt.Println("can nod")
}

func (g *SingGopher) Wink() {
	fmt.Println("wink also can")
}

// START PLAY OMIT
func main() {

	var singGopher *SingGopher
	var gopher Gopher
	var gopher2 Gopher = singGopher

	fmt.Println("values:")
	fmt.Println("singGopher: ", singGopher, singGopher == nil)
	fmt.Println("gopher: ", gopher, gopher == nil)
	fmt.Println("gopher2: ", gopher2, gopher2 == nil)

	fmt.Println("types: ")
	fmt.Printf("singGopher: %#v\n", singGopher)
	fmt.Printf("gopher: %#v\n", gopher)
	fmt.Printf("gopher2: %#v\n", gopher2)

}

// END PLAY OMIT
