package main

import "fmt"

func main() {
	// en el corchete se pone el tipo de dato de la llave y se le indica el tipo de dato que tendrá cada uno de sus valores
	newMap := make(map[string]int)
	newMap["Mauro"] = 23
	newMap["Lucy"] = 25
	newMap["Daniel"] = 25

	for key, value := range newMap {
		fmt.Println(key, value)
	}

	fmt.Println(newMap)

	// isOk sirve para saber si el key existe o no en el mapa
	// devolviendo true si existe o false si no existe
	value, isOk := newMap["Mauro"]
	fmt.Println(value, isOk)
}
