package main

import "fmt"

var secretNumber int

func main() {

	//classic for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//undefined
	for {
		fmt.Println("Ingrese el numero secreto: ")
		fmt.Scanln(&secretNumber)
		if secretNumber == 100 {
			fmt.Println("Felicidades lo has encontrado")
			break
		}
		fmt.Println("Continue")
	}

	//goto
	var i int = 0
	//start coding place label rutina
RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("voy a rutina sumando 2 + 1")
			goto RUTINA
		}
		fmt.Println("Valor de i: ", i)
		i++
	}

}
