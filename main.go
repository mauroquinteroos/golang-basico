package main

import (
	"fmt"
	// "golang_basico/mypackage"
	pk "golang_basico/mypackage"
)

// type sirve para indicar tipo de dato
type car struct {
	brand string
	year  int
}

func main() {
	// Instanciar una estructura vacía
	var myCar1 car
	myCar1.brand = "Ford"
	fmt.Println(myCar1)

	// Instanciar una estructura de manera corta
	myCar2 := car{brand: "Tesla", year: 2021}
	fmt.Println(myCar2)

	// var myCar3 = mypackage.CarPublic{Brand: "Ferrari", Year: 2021}
	var myCar3 = pk.CarPublic{Brand: "Ferrari", Year: 2021}
	fmt.Println(myCar3)

	// var myCar4 = pk.carPrivate{Brand: "Ferrari", Year: 2021} // (al ser privado no se puede acceder)
	pk.PrintMessage()
}
