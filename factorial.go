package main

import "fmt"

func Factorial() {

	fmt.Println("\n=== MÓDULO DE Factorial ===")

	// Aquí comenzará el programa

	leervalores()

}

func leervalores() int {

	var n int
	fmt.Print("\nHasta donde desea sumar : ")
	fmt.Scan(&n)
	bucle(n)
	return n

}

func bucle(n int) {

	var sum, i int

	for i = 1; i <= n; i++ {
		sum += i
		fmt.Printf("%d + ", i)
	}
	//fmt.Print("\nLa Suma de los numeros es: ", sum)

	fmt.Printf("\nLa suma desde 1 hasta %d es: %d", n, sum)
}
