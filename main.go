package main

import "fmt"

func main() {

	pantallaprincipal()

}

func pantallaprincipal() {
	fmt.Println("\nPROGRAMA GENERAL DE MATEMATICAS")
	fmt.Printf("\n 1 - Calculo de Ecuaciones Cuadraticas: ")
	fmt.Printf("\n 2 - Calculo de Hipotenusa : ")
	fmt.Printf("\n 3 - Calculo de Factorial : ")
	fmt.Printf("\n 4 - Calculo del Maximo Comun Divisor: ")
	fmt.Printf("\n 5 - Resolver Sistemas de n Incognitas : ")
	fmt.Printf("\n 6 - Programa para Calcular Precio de Compra : ")
	fmt.Printf("\nPara Ingresar indique el numero de Modulo: ")
	leeropcion()
}

func leeropcion() {

	var opcion int
	fmt.Scan(&opcion)
	modulomenu(opcion)
	//return opcion

}

func modulomenu(opcion int) {
	switch opcion {

	case 1:
		moduloEcuacion()
	case 2:
		moduloHyp()

	case 3:
		Factorial()

	case 4:

		mcd2()

	case 5:
		ecuacionesCramer()

	case 6:

		Venta()

	default:
		fmt.Print("Opcion incorrecta")

	}
}
