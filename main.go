package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	text = strings.ToLower(text)
	var textReverse string = ""

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	fmt.Println(textReverse)
	if text == textReverse {
		fmt.Println("Es palindrome")
	} else {
		fmt.Println("No es palindrome")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindrome("Ama")
}
