package mypackage

import "fmt"

// CarPublic: Estructura Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprime un mensaje
func PrintMessage() {
	fmt.Println("Hello world")
}

func printPrivateMessage() {
	fmt.Println("Hello world")
}
