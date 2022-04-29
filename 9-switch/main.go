package main

import "fmt"

func main() {
	fmt.Printf(`Opciones
		1. Registrarse
		2. Ingresar
		3. Pedir ayuda
	`)
	fmt.Print("Ingrese opcion: ")

	var opcion int
	_, _ = fmt.Scanf("%d", &opcion)

	// No necesita break
	// Se permite matching
	switch {
	case opcion == 1:
		fmt.Println("Vamos a registrar...")
	case opcion == 2:
		fmt.Println("Vamos a ingresar...")
	case opcion == 3:
		fmt.Println("Vamos a darte ayuda...")
	case opcion < 10:
		fmt.Println("La opcion es menor a 10, encontraste una opcion secreeeta")
	default:
		fmt.Println("No entendi...")
	}
}
