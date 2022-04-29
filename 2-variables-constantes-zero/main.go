package main

import "fmt"

func main() {
	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Variables
	base := 12
	var altura int = 14
	var area int

	area = base * altura
	fmt.Println("area", area)

	// Zero values
	// Valores al momento de inicializar
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("default int ", a)
	fmt.Println("default float ", b)
	fmt.Println("default string: ", c)
	fmt.Println("default bool ", d)
}
