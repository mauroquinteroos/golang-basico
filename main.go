package main

import (
	"fmt"
	pn "golang_basico/mypackage"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

// Receiver functions with Value receiver
func (myPc pc) ping() {
	myPc.brand = "Apple"
	fmt.Println("local object: ", myPc)
}

// Receiver functions with Pointer receiver
func (myPc *pc) duplicateRam() {
	myPc.disk *= 2
	fmt.Println("local object: ", myPc)
}

func main() {
	a := 50
	// b será el puntera de a (la dirección del espacio en memoria donde se guarda a)
	b := &a

	fmt.Println("a:", a)
	fmt.Printf("%p\n", &a)
	fmt.Println("b: ", b)

	fmt.Println("--------------------------------")

	*b = 100
	fmt.Println("b: ", *b)
	fmt.Println("a:", a)

	fmt.Println("--------------------------------")

	myPc := pc{ram: 4, disk: 500, brand: "Dell"}
	myPc.ping()
	fmt.Println("original object: ", myPc)

	fmt.Println("--------------------------------")

	myPc.duplicateRam()
	fmt.Println("original object: ", myPc)

	fmt.Println("--------------------------------")

	person1 := pn.Person{Name: "Mauro", Doc: 77382771, Age: 23}
	fmt.Println("person1: ", person1.Print())
	person1.AddAge()
	fmt.Println("person1: ", person1.Print())
}
