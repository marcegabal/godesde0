package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["Mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises["Argentina"])

	//otra forma de crear map
	campeonato := map[string]int{
		"Barcelona":39,
		"Real Madrid":38,
		"River": 99,
		"Chivas": 3,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato{
		fmt.Printf("Equipo %s, campeonatos %d \n", equipo, puntaje)
	}

	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("Puntaje: %d, existe: %t\n", puntaje, existe)
}
