//Versión 1.0

// Eliminación de Gauss.
//Sustitución hacia atrás.
//Requiere pivotes distintos de cero.

package main

import "fmt"

func ecuacionesGauss() {

	fmt.Println("\n=== MÓDULO para Solucion de Sistemas de Ecuaciones===")
	leern()
}

func leern() {

	var n int

	fmt.Print("\nIngrese el valor de n: ")
	fmt.Scan(&n)

	matriz := prepararMatriz(n)
	leerCoeficientes(matriz, n)
	imprimirMatriz(matriz, n)
	resolverGauss(matriz, n)
	imprimirMatrizGauss(matriz, n)
	soluciones := sustitucionAtras(matriz, n)
	imprimirSoluciones(soluciones, n)
	pgmReporte(matriz, soluciones, n)

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

	fmt.Println("\nMatriz aumentada:")

	for i := 0; i < n; i++ {

		for j := 0; j < n+1; j++ {

			fmt.Printf("%8.2f", matriz[i][j])

		}

		fmt.Println()

	}
}

func resolverGauss(matriz [][]float64, n int) {

	recorrerPivotes(matriz, n)

}

func recorrerPivotes(matriz [][]float64, n int) {

	for pivote := 0; pivote < n-1; pivote++ {

		procesarFilas(matriz, n, pivote)

	}

}

func procesarFilas(matriz [][]float64, n int, pivote int) {

	for fila := pivote + 1; fila < n; fila++ {

		calcularFactor(matriz, pivote, fila)

	}

}

func calcularFactor(matriz [][]float64, pivote int, fila int) {

	pivoteActual := matriz[pivote][pivote]

	elemento := matriz[fila][pivote]

	factor := elemento / pivoteActual

	actualizarFila(matriz, pivote, fila, factor)

}

func actualizarFila(matriz [][]float64, pivote int, fila int, factor float64) {

	for columna := pivote; columna < len(matriz[0]); columna++ {

		matriz[fila][columna] =
			matriz[fila][columna] -
				factor*matriz[pivote][columna]

	}

}

func sustitucionAtras(matriz [][]float64, n int) []float64 {

	// Crear el vector donde se almacenarán las soluciones
	soluciones := make([]float64, n)

	// Recorrer las filas desde la última hasta la primera
	for fila := n - 1; fila >= 0; fila-- {

		calcularVariable(matriz, soluciones, fila, n)

	}

	return soluciones

}

func calcularVariable(matriz [][]float64, soluciones []float64, fila int, n int) {

	// Comenzamos con el resultado de la ecuación
	resultado := matriz[fila][n]

	// Restar las variables que ya fueron calculadas
	for columna := fila + 1; columna < n; columna++ {

		resultado = resultado - matriz[fila][columna]*soluciones[columna]

	}

	// Dividir entre el coeficiente del pivote
	soluciones[fila] = resultado / matriz[fila][fila]

}

func imprimirSoluciones(soluciones []float64, n int) {

	fmt.Println("\nSolución del sistema:\n")

	for i := 0; i < n; i++ {

		fmt.Printf("X%d = %.1f\n", i+1, soluciones[i])

	}

}

func imprimirMatrizGauss(matriz [][]float64, n int) {

	fmt.Println("\nMatriz después de la Eliminación de Gauss:\n")

	for i := 0; i < n; i++ {

		for j := 0; j < n+1; j++ {

			fmt.Printf("%10.2f", matriz[i][j])

		}

		fmt.Println()

	}

}
