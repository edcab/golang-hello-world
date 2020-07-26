package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//Ejecuta de forma asincrona
	go miNombreLentooooo("Edwin")
	fmt.Println("Estoy aqui")
	var x string
	//En el runtime de go no espera que la funcion asincrona termine
	fmt.Scanln(&x)
}

func miNombreLentooooo(nombre string) {
	letras := strings.Split(nombre, "")

	for _, eachChar := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(eachChar)
	}
}
