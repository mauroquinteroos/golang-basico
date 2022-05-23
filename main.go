package main

import "fmt"

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func main() {
	cuadrado1 := cuadrado{base: 5}
	fmt.Println("Area: ", cuadrado1.area())

	rectangulo1 := rectangulo{base: 5, altura: 10}
	fmt.Println("Area: ", rectangulo1.area())
}
