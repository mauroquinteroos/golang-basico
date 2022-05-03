package main

import "fmt"

func main() {
	// defer ejecuta una línea de código antes de que se termine de ejecutar el programa
	// Se utiliza para cerrar una conexión a una base de datos, para cerrar un archivo.
	defer fmt.Println("Hello")
	fmt.Println("World")

	for i := 0; i < 10; i++ {
		fmt.Println("Number", i)

		if i == 2 {
			fmt.Println("Number is odd")
			continue // Se salta el resto de la línea y no detiene el ciclo for
		}

		if i == 8 {
			fmt.Println("Number is 8")
			break // Se detiene el ciclo for
		}
	}
}
