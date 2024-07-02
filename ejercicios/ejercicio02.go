package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error
var texto string

func IngresoMultiplicacion() string {
	scanner := bufio.NewScanner(os.Stdin) //recibe un param que es de donde recibe el dato

	for {
		fmt.Println("Numero a multiplicar: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				//panic("Error ---> " + err.Error())
				continue
			} else {

				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		//fmt.Println("--> ", num*i)
		texto += fmt.Sprintf("%d x %d = %d \n", num, i, i*num)
	}

	return texto
}
