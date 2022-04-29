package main

import "fmt"

func main() {
	var a int
	// int (depende del OS)
	// int8, int16, int32, int64
	a = 1
	fmt.Println("int8, int16, int32, int64", a)

	var b uint
	// depende del OS
	// uint8, uint16, uint32, uint64
	b = 0
	fmt.Println("uint8, uint16, uint32, uint64", b)

	var f float32 = 1.1
	// tambien esta float64
	fmt.Println("float32, float64", f)

	// var cadena string = "comillas dobles"
	// var booleano bool

	var complejo complex64 = 10 + 8i
	var complejo128 complex128 = 10 + 8i
	fmt.Println("Complejo 64", complejo)
	fmt.Println("Complejo 128", complejo128)
}
