package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es primer mensaje")
	defer fmt.Println("Este es mensaje final")
	fmt.Println("Este es segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio err que genero panic \n%v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro 1")
	}
}
