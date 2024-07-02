package main

import (
	"fmt"
	"runtime"

	"github.com/marcegabal/godesde0/ejercicios"
	"github.com/marcegabal/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1558)
	fmt.Println(estado)
	fmt.Println(texto)

	//os := runtime.GOOS // puedo poner la asignacion en el if tambien
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("No es Win")
	} else {
		fmt.Println("Esto es Win", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Estto es Linux")
	case "darwin":
		fmt.Println("Estto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	///////////
	fmt.Println("==========EJERCICIO======")
	numej, textej := ejercicios.EjercicioReturn("199")

	fmt.Println(numej)
	fmt.Println(textej)

}
