package main

import (
	"fmt"

	"github.com/marcegabal/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1558)
	fmt.Println(estado)
	fmt.Println(texto)
}
