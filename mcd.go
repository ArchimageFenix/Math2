package main

import "fmt"

func mcd() {

	leernumeros()

}

func leernumeros() {

	var x, y int
	fmt.Println("\n=== MÓDULO para Encontrar el M.C.D o Maximo Comun Divisor===")
	fmt.Println("Primer Valor de A: ")
	fmt.Scan(&x)
	fmt.Println("Primer Valor de B: ")
	fmt.Scan(&y)
	calculomcd(x, y)

}

func calculomcd(x int, y int) { //FUNCION PARA CALCULAR E IMPRIMIR EL MCD USANDO RESTAS SUCESIVAS

	for x != y {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	fmt.Println("El MCD es: ", x)
}
