package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == valor2 {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}

	// With AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// With OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// Convertir texto a número
	value, err := strconv.Atoi("10")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}

	fmt.Println(10 % 2)
}
