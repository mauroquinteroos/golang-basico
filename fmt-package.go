package main

import "fmt"

func main() {
	wordMessage := "Hello World"

	// Println
	fmt.Println(wordMessage)

	name := "Platzi"
	courses := 500

	// Printf
	// %s = string
	// %d = integer
	// En caso no sabemos el tipo de dato se usa %v
	fmt.Printf("%s tiene curso %d cursos\n", name, courses)
	fmt.Printf("%v tiene curso %v cursos\n", name, courses)

	// Sprintf
	message := fmt.Sprintf("%v tiene curso %v cursos", name, courses)
	fmt.Println(message)

	// Tipo de datos de una variable
	fmt.Printf("El tipo de dato de la variable name es: %T\n", name)
	fmt.Printf("El tipo de dato de la variable courses es: %T\n", courses)
}
