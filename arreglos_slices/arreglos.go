package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{7, 66, 99} //puedo hacerlo sin esto

var matriz [20][30]int

func MuestroArreglos() {
	tabla[9] = 33
	tabla[7] = 54

	fmt.Println(tabla)
	tabla2 := [10]string{"P", "J"}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15

	fmt.Println(matriz)
}
