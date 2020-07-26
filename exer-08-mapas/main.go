package main

import "fmt"

func main() {
	//crear un map com make de tipo string y el valor que le vamos asignar tambien va ser string
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	// map de clave tipo string, valor entero
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  28,
		"Chivas":       27,
		"Boca Juniors": 30,
	}

	for _, value := range campeonato {
		fmt.Println(value)
	}

	fmt.Println("Adding...")
	campeonato["River Plate"] = 25
	for key, value := range campeonato {
		fmt.Println(key, value)
	}

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("el Equipo chivas existe ?? %t tiene un puntaje de %d", existe, puntaje)

}
