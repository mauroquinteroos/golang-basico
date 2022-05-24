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

type Figura interface {
	area() float64
}

func calcular(figura Figura) {
	fmt.Println("Area: ", figura.area())
}

func main() {
	cuadrado1 := cuadrado{base: 5}
	rectangulo1 := rectangulo{base: 5, altura: 10}

	calcular(cuadrado1)
	calcular(rectangulo1)

	fmt.Println("--------------------------------")

	myInterface := []interface{}{"Hola", 12, true, 5.9}
	fmt.Println(myInterface...)
}
