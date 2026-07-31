package main

import "fmt"

func ecuacionesCramer() {

	fmt.Println("\n=== MÓDULO para Solucion de Sistemas de Ecuaciones===")
	leern()
}

func leern() {

	var n int

	fmt.Println("Ingrese el valor de n: ")
	fmt.Scan(&n)
	prepararMatriz(n)

}

func prepararMatriz(n int) [][]float64 {
	// Crear las filas
	matriz := make([][]float64, n)

	// Crear las columnas de cada fila
	for i := 0; i < n; i++ { //[][]  crear un slice cuyos elementos también serán slices.

		matriz[i] = make([]float64, n)

	}

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {

			fmt.Printf("%8.2f", matriz[i][j])

		}

		fmt.Println()

	}

	return matriz

}
