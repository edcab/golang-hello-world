package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error", err)
	} else {
		fmt.Println(string(archivo))
	}

}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt") // no lee nada, lo que hace es abrir el archivo y tenerlo listo para lectura.
	if err != nil {
		fmt.Println("Hubo un error", err)
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}

}
