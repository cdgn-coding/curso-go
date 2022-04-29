package main

import "fmt"

func main() {
	// Defer
	// Ejecuta antes de morir la funcion.
	// Utiliza una pila de llamadas defer...
	// Last in first out
	defer fmt.Println("Hola")
	defer fmt.Println("Mundo")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 6 {
			break
		}

		fmt.Printf("i = %d\n", i)
	}
}
