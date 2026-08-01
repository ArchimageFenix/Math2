package main

import "fmt"

func ecuacionesCramer() {

	fmt.Println("\n=== MÓDULO para Solucion de Sistemas de Ecuaciones===")
	leern()
}

func leern() {

	var n int

	fmt.Println("\nIngrese el valor de n: ")
	fmt.Scan(&n)

	matriz := prepararMatriz(n)

	leerCoeficientes(matriz, n)
	imprimirMatriz(matriz, n)

}

func prepararMatriz(n int) [][]float64 {
	// Crear las filas
	matriz := make([][]float64, n)

	// Crear las columnas de cada fila
	for i := 0; i < n; i++ { //[][]  crear un slice cuyos elementos también serán slices.

		matriz[i] = make([]float64, n+1)

	}

	return matriz

}

func leerCoeficientes(matriz [][]float64, n int) {

	for i := 0; i < n; i++ {

		fmt.Printf("\nEcuación %d\n", i+1)

		for j := 0; j < n+1; j++ {

			if j == n {
				fmt.Print("Resultado: ")
			} else {
				fmt.Printf("Coeficiente X%d: ", j+1)
			}

			fmt.Scan(&matriz[i][j])
		}
	}

}

func imprimirMatriz(matriz [][]float64, n int) {

	fmt.Println("\nMatriz aumentada:\n")

	for i := 0; i < n; i++ {

		for j := 0; j < n+1; j++ {

			fmt.Printf("%8.2f", matriz[i][j])

		}

		fmt.Println()

	}
}
