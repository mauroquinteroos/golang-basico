// El nombre de la carpeta en la cual está guardado
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declaracion de constantes
	const pi1 float64 = 3.141592
	const pi2 = 3.1415

	fmt.Println("pi1:", pi1)
	fmt.Println("pi2:", pi2)

	// Declaracion de variable corta
	base := 12 // Go infiere el tipo de dato

	// Declaracion de variable larga, cuando no inicialice la variable
	var altura int = 14
	var area int = 30 // No es una buena práctica declarar y asignar el valor de esta manera

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)

	// Zero values, es el valor por defecto que va tener una variable al no tener un valor.
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// Calcular el area cuadrada
	const baseCuadrado = 10.00
	areaCuadrada := baseCuadrado * baseCuadrado
	fmt.Println("Área cuadrada:", areaCuadrada)
	fmt.Println("Tipo de dato:", reflect.TypeOf((areaCuadrada)))

	// Asignación múltiple
	j, k, l := 1, 2.6, "Mauro"
	fmt.Println("j:", j)
	fmt.Println("k:", k)
	fmt.Println("l:", l)
}
