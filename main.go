package main

import "fmt"

func updateArray(array [4]int) [4]int {
	array[0] = 100
	return array
}

func updateSlice(slice []int) []int {
	slice[0] = 100
	return slice
}

func main() {
	var array1 [4]int
	array1[0] = 1
	array1[1] = 2
	fmt.Println(array1)
	fmt.Println("len: ", len(array1))
	fmt.Println("cap: ", cap(array1))
	fmt.Printf("El tipo de dato es: %T\n", array1)
	fmt.Printf("El espacio en memoria es: %p\n", &array1)

	fmt.Println("----------------------------------------------------------------")

	arrayTransform := updateArray(array1)
	fmt.Println(arrayTransform)
	fmt.Printf("El espacio en memoria del arrayTransform es: %p\n", &arrayTransform)
	fmt.Println(array1)
	fmt.Printf("El espacio en memoria del array es: %p\n", &array1)

	fmt.Println("----------------------------------------------------------------")

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	fmt.Println("len: ", len(slice1))
	fmt.Println("cap: ", cap(slice1))
	fmt.Printf("El tipo de dato es: %T\n", slice1)
	fmt.Printf("El espacio en memoria es: %p\n", &slice1)

	fmt.Println("----------------------------------------------------------------")

	sliceTransform := updateSlice(slice1)
	fmt.Println(sliceTransform)
	fmt.Printf("El espacio en memoria del sliceTransform es: %p\n", &sliceTransform)
	fmt.Println(slice1)
	fmt.Printf("El espacio en memoria del slice es: %p\n", &slice1)

	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)
	fmt.Println("len: ", len(slice2))
	fmt.Println("cap: ", cap(slice2))

	fmt.Println("----------------------------------------------------------------")

	// Slicing: array[start:end] o slice[start:end]
	// start es inclusivo, si incluye el elemento en esta posición
	// end es exclusivo, no incluye el elemento en esta posición
	slicing := slice1[1:3]
	fmt.Println(slicing)
	fmt.Println("len: ", len(slicing))
	fmt.Println("cap: ", cap(slicing))
	fmt.Println(slice1[1:3])
	fmt.Println(slice1[3:])

	fmt.Println("----------------------------------------------------------------")

	newSlice1 := append(slice1, 7)
	fmt.Println(newSlice1)

	fmt.Println("----------------------------------------------------------------")

	slice3 := []int{8, 9, 10}
	slice3 = append(slice3, newSlice1...)
	fmt.Println(slice3)
}
