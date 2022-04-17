// El nombre de la carpeta en la cual está guardado
package main

import (
	"fmt"
)

func main() {
	// Calcular el area cuadrada
	const baseCuadrado = 10.00
	areaCuadrada := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrada)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("La suma es: ", result)

	// Resta
	result = x - y
	fmt.Println("La resta es: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("La multiplicación es: ", result)

	// Division
	result = y / x
	fmt.Println("La división es: ", result)

	// Modulo
	result = y % x
	fmt.Println("El módulo es: ", result)

	// Incremental
	x++
	fmt.Println("El valor de x es: ", x)

	// Decremental
	x--
	fmt.Println("El valor de x es: ", x)
}
