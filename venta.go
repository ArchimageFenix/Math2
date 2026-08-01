package main

import "fmt"

func Venta() {

	var precio, cantidad float64 //DECLARACION DE VARIABLES FLOAT64

	fmt.Println("\tPrograma de Calculo de Precio de Compra ")
	fmt.Println("Ingrese el Precio del Articulo (C$) ")
	fmt.Scan(&precio)
	fmt.Println("Ingrese el Cantidad de Articulos: (C$)")
	fmt.Scan(&cantidad)
	calcularPreciocompra(precio, cantidad) //FUNCION QUE ENVIA
	//VALORES A

}

func calcularPreciocompra(precio float64, cantidad float64) {

	var precioFinal float64

	precioFinal = precio * cantidad
	fmt.Printf("\nEl Precio a Pagar es de: C$ %.2f", precioFinal)

}
