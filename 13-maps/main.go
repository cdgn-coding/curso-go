package main

import "fmt"

func main() {
	edades := make(map[string]int)
	edades["Jose"] = 25
	edades["Pepito"] = 20
	fmt.Println(edades)

	// Recorrido de maps
	// Se obtienen de forma aleatorea
	for i, valor := range edades {
		fmt.Println(i, "tiene edad", valor)
	}

	// Encontrar un valor
	fmt.Println("jose tiene edad", edades["Jose"])

	// No encontrar un valor
	fmt.Println("si buscamos por un indice que no existe devuelve", edades["no existe"], "el zero value")

	// la inspeccion de mapas devuelve un segundo argumento para verificar
	_, ok := edades["Carlos"]
	fmt.Println("Carlos existe en el mapa?", ok)
}
