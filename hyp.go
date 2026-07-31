package main

import (
	"fmt"
	"math"
)

func moduloHyp() {

	fmt.Println("\n=== MÓDULO para Resolver Triangulos Rectangulos ===")

	// Aquí comenzará el programa

	fmt.Print("\nPrograma para Resolver Triangulos Rectangulos: ")
	leercatetos()

}

//Ejemplo 1: Los catetos de un triángulo rectángulo miden 12 cm y 5 cm. ¿Cuánto mide la
//hipotenusa?
//usando el teorema de pitagoras

func leercatetos() {

	var B, C float64
	fmt.Printf("\nIngrese el Cateto B: ")
	fmt.Scan(&B)
	fmt.Printf("\nIngrese el Cateto C: ")
	fmt.Scan(&C)
	hipotenusa(B, C)

}

func hipotenusa(B float64, C float64) {
	var A, B1, C1 float64
	B1 = B * B
	C1 = C * C
	fmt.Printf("\nEl Cateto B Es: %.2f", B1)
	fmt.Printf("\nEl Cateto C Es: %.2f ", C1)
	A = math.Sqrt(B1 + C1)
	fmt.Printf("\nLa Hipotenusa Es: %.2f ", A)

}
