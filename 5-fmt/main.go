package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)
	// Agrega salto de linea
	fmt.Println(helloMessage, worldMessage)

	fmt.Printf("Cadena es %s, numbero es %d, cualquier tipo %v\n", "Curso", 25, "25")

	cadena := fmt.Sprintf("Contenido de la cadena formateado. Numero == %d", 25)
	fmt.Println(cadena)

	fmt.Printf("mostrar tipo de variable. El tipo de helloMessage es %T\n", helloMessage)
}
