package main

import (
	"fmt"
	"time"
)

func main() {
	// Un canal es un espacio de memoria en el que van a dialogar las go routines
	// el espacio de memoria debe tener un tipo de dato

	canal1 := make(chan time.Duration) // chan tipo time.Duration

	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")

	//Estoy escuchando el canal1 porfavor no termine la ejecución
	//le digo que espere hasta que canal1 reciba un valor
	msg := <-canal1
	fmt.Println(msg)

}

func bucle(canal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 10000000000; i++ {

	}

	final := time.Now()

	canal <- final.Sub(inicio) // .Sub() es la funcion que me retorna la duración

}
