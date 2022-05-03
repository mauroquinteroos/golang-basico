package main

import "fmt"

func main() {
	modulo2 := 10 % 2

	// 1° forma (con condición determinada)
	switch modulo2 {
	case 0:
		fmt.Println("El número es par")
	case 1:
		fmt.Println("El número es impar")
	default:
		fmt.Println("No se puede calcular")
	}

	// 2° forma (con condición determinada)
	switch modulo1 := 10 % 2; modulo1 {
	case 0:
		fmt.Println("El número es par")
	case 1:
		fmt.Println("El número es impar")
	default:
		fmt.Println("No se puede calcular")
	}

	// 3° forma (sin condición determinada)
	value := 80
	switch {
	case value > 100:
		fmt.Println("El valor es mayor a 100")
	case value < 0:
		fmt.Println("El valor es menor a 0")
	default:
		fmt.Println("El valor no cumple ninguna condicion")
	}

}
