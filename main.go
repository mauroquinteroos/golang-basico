package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

// func getResult(x int, y int, z int, message string) {
func getResult(x, y, z int, message string) {
	fmt.Println(message)
	fmt.Println(x + y + z)
}

func calculateOperation(x, y, z int) int {
	return x * y * z
}

func calculateSumRest(x, y, z int) (a, b int) {
	a = x + y + z
	b = x - y - z
	return a, b
}

func main() {
	printMessage("Hello World")
	printMessage("Mauro Quinteros")

	getResult(1, 2, 3, "Resultado")

	value1 := calculateOperation(1, 2, 3)
	fmt.Println(value1)

	value2, value3 := calculateSumRest(1, 2, 3)
	value4, _ := calculateSumRest(1, 2, 3)
	fmt.Println(value2, value3)
	fmt.Println(value4)
}
