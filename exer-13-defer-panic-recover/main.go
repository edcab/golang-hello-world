package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	ejemploPanic()
}

//panic, es para forzar un error
func ejemploPanic() {
	//Implementamos defer para encontrar si hubo error o no
	defer func() {
		//Si encontro un panic graba el resultado del ultimo panic,
		//si no hubo panic entonces recover devuelve nil
		reco := recover()
		if reco != nil {
			log.Fatalf("Hubo un error, es el siguiente: \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}

//defer es una instruccion que se va a eejcutar si o si
//cuando se detecta que se sale de la funcion
func ejemploDefer() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	defer f.Close()
	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}
}
