package main

import "fmt"

func main() {
	estado := true

	if estado = false; estado == true {
		fmt.Println("Estado: ", estado)
	} else {
		fmt.Println("Estado: ", estado)
	}

	num := 6
	switch num {
	case 5:
		fmt.Println("El valor de num es cinco ? ", num)
	case 4:
		fmt.Println("El valor de num es cinco ? ", num)
	case 3:
		fmt.Println("El valor de num es cinco ? ", num)
	case 2:
		fmt.Println("El valor de num es cinco ? ", num)
	case 1:
		fmt.Println("El valor de num es cinco ? ", num)
	default:
		fmt.Println("El valor de num no esta entre 1 y 5.")
	}
}
