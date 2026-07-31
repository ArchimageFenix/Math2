package main

import (
	"fmt"
	"math"
)

func moduloEcuacion() {

	fmt.Println("\n=== MÓDULO DE ECUACIONES CUADRÁTICAS ===")

	// Aquí comenzará el programa

	fmt.Println("\nPrograma de Resolucionde EC de Segundo Grado")
	fmt.Println("\nPor medio de la Formula General")
	leerdatos() //LEE LOS VALORES DE A, B Y C

}

//FUNCION PARA LEER DATOS QUE ENVIARA A MAIN

func leerdatos() (float64, float64, float64) {

	var a, b, c float64
	fmt.Print("\nIngrese el Valor de A : ")
	fmt.Scan(&a)
	fmt.Print("\nIngrese el Valor de B : ")
	fmt.Scan(&b)
	fmt.Print("\nIngrese el Valor de C : ")
	fmt.Scan(&c)
	calculardet(a, b, c)
	return a, b, c

}

// FUNCION SEGUNDA DE CALCULO DE DETERMINANTE
func calculardet(a float64, b float64, c float64) {

	if a != 0 { //VALIDACION DE VALOR A
		det := ((b * b) - (4 * a * c))
		validardet(a, b, c, det)
		//return det
	} else {

		fmt.Printf("\nERROR el valor de A no puede ser cero, La Ecuacion no es Cuadratica!!!")

	}

}

// FUNCION TRCERA PARA VAERIFICAR SI ES NEGATIVO EL DET
func validardet(a float64, b float64, c float64, det float64) float64 {

	if det >= 0 {
		//fmt.Print("El determinante Positivo es: ", det)
		calcularx(a, b, c, det)

	} else {
		fmt.Println("Valor de D: ", det)
		fmt.Print("\nSistema sin solucion real: ")
	}

	return det

}

func calcularx(a float64, b float64, c float64, det float64) {

	var cociente float64
	cociente = 2 * a
	x1 := ((b * -1) + math.Sqrt(det)) / (cociente)
	fmt.Println(" Valor de A: ", a)
	fmt.Println(" Valor de B: ", b)
	fmt.Println(" Valor de C: ", c)
	fmt.Println(" Valor de D: ", det)
	fmt.Printf("\nEl valor de x1 es: %.2f", x1)
	x2 := ((b * -1) - math.Sqrt(det)) / (cociente)
	fmt.Printf("\nValor de x2 es: %.2f", x2)

}
