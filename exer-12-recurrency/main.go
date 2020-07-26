package main

import "fmt"

func main() {
	exponencia(1)
}

func exponencia(numero int64) {
	if numero > 1000000000000000000 {
		return
	}

	fmt.Println(numero)

	exponencia(numero * 2)
}
