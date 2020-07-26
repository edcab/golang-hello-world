package main

import (
	"fmt"
	"strconv"
)

var num int
var text string
var status bool

func main() {
	num2, num3, num4, myString := 2, 3, 4, "Edwin"

	var moneda int64 = 1000

	myString = strconv.Itoa(int(moneda))

	fmt.Println("Hola desde exer-02")
	fmt.Println("Valor en num2: ", num2)
	fmt.Println("Valor en num3: ", num3)
	fmt.Println("Valor en num4: ", num4)
	fmt.Println("Valor en myString: ", myString)
}
