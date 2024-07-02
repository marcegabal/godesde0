package ejercicios

import "strconv"

func EjercicioReturn(valor string) (int, string) {
	var texto string
	// string to int
	num, err := strconv.Atoi(valor)
	if err != nil {
		panic(err)
	}

	if num > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}

	return num, texto
}
