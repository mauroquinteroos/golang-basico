package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 4)
	c <- "hello"
	c <- "world"
	fmt.Println("len: ", len(c))
	fmt.Println("cap: ", cap(c))

	close(c)

	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("message 1", email1)
	go message("message 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-email1:
			fmt.Println("Email recibido de email1: ", msg1)
		case msg2 := <-email2:
			fmt.Println("Email recibido de email2: ", msg2)
		}
	}
}
