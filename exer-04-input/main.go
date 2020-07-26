package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var leyenda string

func main() {

	fmt.Println("Ingrese Número 1: ")
	fmt.Scanf("%d", &numero1)

	fmt.Println("Número ingresado: ", numero1)

	fmt.Println("Ingrese Número 2: ")
	fmt.Scanf("%d", &numero2)

	fmt.Println("Número ingresado: ", numero2)

	//bufio
	fmt.Println("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado := numero1 + numero2
	fmt.Println(leyenda, resultado)

}
