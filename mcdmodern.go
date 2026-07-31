package main

import "fmt"

func mcd2() {
	fmt.Println("\n=== MÓDULO para Encontrar el M.C.D o Maximo Comun Divisor===")
	leernumeros2()

}

func leernumeros2() {

	var x, y int

	fmt.Println("\nValor de A: ")
	fmt.Scan(&x)
	fmt.Println("\nValor de B: ")
	fmt.Scan(&y)
	calculomcd2(x, y)

}

func calculomcd2(x int, y int) { //FUNCION PARA CALCULAR E IMPRIMIR EL MCD USANDO RESTAS SUCESIVAS

	for y != 0 {

		fmt.Printf("A = %d, B = %d\n", x, y)
		fmt.Printf("%d %% %d = %d\n", x, y, x%y)
		x, y = y, x%y
	}

	fmt.Println("El MCD es: ", x)

}
