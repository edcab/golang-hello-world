package main

import "fmt"

//las funciones son tambien tipos de datos
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	// funciones anonimas son funciones que no tienen nombre
	// ejemplo: sirven para crear operaciones
	fmt.Printf("Sumando 5 + 7: %d", Calculo(5, 7))

	//Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Restamos 5 - 7: %d", Calculo(5, 7))

	Operaciones()

	tabladel := 2
	miTabla := Tabla(tabladel)

	for i := 0; i < 11; i++ {
		fmt.Println(miTabla())
	}

}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13

		return a + b
	}

	fmt.Println(" Sumando en operaciones: ", resultado())
}

//Closures
//Creamos una funcion tabla que recibe un entero y a su vez retorna una funcion que retorna un entero
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
