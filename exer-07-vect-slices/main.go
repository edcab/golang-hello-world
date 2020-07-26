package main

import "fmt"

func main() {
	//array
	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, num := range tabla {
		fmt.Println(num)
	}

	//matriz
	var matriz [5][7]int
	matriz[3][5] = 1
	fmt.Println(matriz)

	//slices, son vectores dinamicos, vectores donde se pueden ampliar las dimensiones en tiempo de ejecución
	variante2()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}

	porcion := elementos[3:] //creamos un slice desde la posicion 3 de elementos hasta la ultima

	fmt.Println(porcion)

	variante3()
}

func variante3() {
	//array, cantidad, capacidad(inicialmente crea un vector con cinco elementos con una capacidad preparada de hasta 20)
	elementos := make([]int, 5, 20)
	fmt.Printf("Length %d, Capacidad %d.", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0) // crear un slice vacio

	//crea un elemento en nums
	for i := 0; i < 100; i++ {
		nums = append(nums, i) // append lo que recibe es un slice y devuelve un slice
	}

	//nos damos cuenta que agrego elementos y redimensiona automaticamente
	fmt.Printf("\n redimensionando: Length %d, Capacidad %d", len(nums), cap(nums))
}
