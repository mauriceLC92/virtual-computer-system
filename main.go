package main

import (
	"fmt"

	"github.com/mauriceLC92/virtual-computer-system/gmachine"
)

func main() {
	fmt.Println("Gmachine!")
	g := gmachine.New()

	g.RunProgram([]uint64{
		uint64(gmachine.SETA), 5,
		uint64(gmachine.HALT),
	})
}
