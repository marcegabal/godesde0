package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin) //recibe un param que es de donde recibe el dato

	fmt.Println("Ingrese num 1: ")
	//este if es para que siga solo si el usuario tipea algo
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error ---> " + err.Error())
		}
	}

	fmt.Println("Ingrese num 2: ")
	//este if es para que siga solo si el usuario tipea algo
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error ---> " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	//este if es para que siga solo si el usuario tipea algo
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
