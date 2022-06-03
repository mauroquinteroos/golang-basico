package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // sirve para liberar la goroutine del waitgroup
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")

	wg.Add(1) // Agrega una goroutine al WaitGroup para que espere la ejecución antes de que la goroutine main termine
	go say("World", &wg)

	wg.Add(1)
	// Función anónima
	go func(text string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(text)
	}("Bye!", &wg)

	wg.Wait() // Sirve para decirle a la goroutine principal que espere a las demás goroutines
}
