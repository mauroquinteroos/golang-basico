package main

import "fmt"

func main() {
	// ciclo for
	for i := 0; i < 10; i++ {
		fmt.Println("ciclo for: ", i)
	}

	fmt.Printf("\n")

	// ciclo while
	i := 0
	for i < 10 {
		fmt.Println("ciclo while: ", i)
		i++
	}

	fmt.Printf("\n")

	// for forever
	counterForever := 0
	for {
		fmt.Println("ciclo for forever: ", counterForever)
		counterForever++
	}
}
