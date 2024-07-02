package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marcegabal/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

// aca se crea pero se pisa todo el contenido
func GrabarTabla() {
	var texto string = ejercicios.IngresoMultiplicacion()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error en crear archivo " + err.Error())
		return //para salir del metodo
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

// aca concateno enl contenido
func SumarTabla() {
	var texto string = ejercicios.IngresoMultiplicacion()
	if !Append(fileName, texto) {
		fmt.Println("Error concatenar")
	}
}

func Append(filen string, txt string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error append archivo al abrir " + err.Error())
		return false
	}

	_, err = arch.WriteString(txt)
	if err != nil {
		fmt.Println("Error append archivo " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error leyendo archivo al " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)
	}
	archivo.Close()
}
