package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(1)
	fmt.Println(numero, estado)

	fmt.Println(tres(3))

	cuatro(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func tres(numero int) (z int) {
	z = numero * 2
	return z

}

//Parametros variables y uso de range
func cuatro(numero ...int) {
	total := 0
	//devuelve dos valores, numero del elemento y el valor del array numero, en este caso _, num
	for _, num := range numero {
		total = total + num
	}

	fmt.Println("La suma es: ", total)
}
