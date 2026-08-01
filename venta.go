package main

import "fmt"

func Venta() {

	var precio, cantidad float64 //DECLARACION DE VARIABLES FLOAT64

	fmt.Println("\nPrograma de Calculo de Precio de Compra ")
	fmt.Println("\nIngrese el Precio del Articulo (C$) ")
	fmt.Scan(&precio)
	fmt.Println("\nIngrese el Cantidad de Articulos: (C$)")
	fmt.Scan(&cantidad)
	calcularPreciocompra(precio, cantidad) //FUNCION QUE ENVIA
	//VALORES A LA FUNCION calcularPreciocompra

}

func calcularPreciocompra(precio float64, cantidad float64) {
	//ESTA FUNCION OBTIENE DOS PARAMETROS TIPO FLOAT QUE RCIBE DE MAIN
	// EL OBJETIVO ES QUE ESTA CALCULE EL PRECIO QUE ELCLIENTE DEBE PAGAR
	//EN BASE A LA CANTIDAD DE ARTICULOS DEJANDO POR TANTO A MAIN COMO
	//LECTOR DE DATOS SOLAMENTE
	var precioFinal float64 //DECLARACION DE VARIABLE

	precioFinal = (precio * cantidad)

	fmt.Printf("\nEl Precio a Pagar es de: C$ %.2f", precioFinal)

}
