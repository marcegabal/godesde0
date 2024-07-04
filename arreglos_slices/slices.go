package arreglos_slices

import "fmt"

var tablaS []int = []int{2,8}

var arreglo [10]int = [10]int{6, 88, 99, 2, 9, 73, 21, 45, 66, 79}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:] //desde un vector de pos 3 hasta final
	porcion2 := arreglo[:5] //desde principio hasta donde le digo
	porcion3 := arreglo[6:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad(){
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++{
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}