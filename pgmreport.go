package main

import (
	"fmt"
	"os"
	"time"
)

func pgmReporte(matriz [][]float64, soluciones []float64, n int) {

	// Fecha y hora actual
	ahora := time.Now()

	// Nombre del archivo
	nombre := fmt.Sprintf(
		"PGM_Gauss_%s.txt",
		ahora.Format("20060102_150405"),
	)

	// Crear archivo
	archivo, err := os.Create(nombre)

	if err != nil {
		fmt.Println("\nError al crear el reporte.")
		return
	}

	defer archivo.Close()

	// Encabezado
	fmt.Fprintln(archivo, "===========================================")
	fmt.Fprintln(archivo, "PROGRAMA GENERAL DE MATEMÁTICAS")
	fmt.Fprintln(archivo, "Módulo: Eliminación de Gauss")
	fmt.Fprintln(archivo, "===========================================")
	fmt.Fprintln(archivo)

	fmt.Fprintf(archivo, "Fecha : %s\n", ahora.Format("02/01/2006"))
	fmt.Fprintf(archivo, "Hora  : %s\n", ahora.Format("15:04:05"))
	fmt.Fprintf(archivo, "Número de incógnitas: %d\n", n)

	fmt.Fprintln(archivo)
	fmt.Fprintln(archivo, "-------------------------------------------")
	fmt.Fprintln(archivo, "MATRIZ TRIANGULAR SUPERIOR")
	fmt.Fprintln(archivo, "-------------------------------------------")
	fmt.Fprintln(archivo)

	// Matriz
	for i := 0; i < n; i++ {

		for j := 0; j < n+1; j++ {

			fmt.Fprintf(archivo, "%12.2f", matriz[i][j])

		}

		fmt.Fprintln(archivo)

	}

	fmt.Fprintln(archivo)
	fmt.Fprintln(archivo, "-------------------------------------------")
	fmt.Fprintln(archivo, "SOLUCIÓN DEL SISTEMA")
	fmt.Fprintln(archivo, "-------------------------------------------")
	fmt.Fprintln(archivo)

	// Soluciones
	for i := 0; i < n; i++ {

		fmt.Fprintf(
			archivo,
			"X%d = %.4f\n",
			i+1,
			soluciones[i],
		)

	}

	fmt.Fprintln(archivo)
	fmt.Fprintln(archivo, "===========================================")
	fmt.Fprintln(archivo, "Fin del reporte")
	fmt.Fprintln(archivo, "===========================================")

	fmt.Println("\nReporte generado correctamente.")
	fmt.Println("Archivo:", nombre)

}
