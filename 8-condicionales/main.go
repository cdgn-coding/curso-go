package main

import "fmt"

func main() {
	var numero int
	var err error
	var n int
	fmt.Print("Ingrese numero=")
	n, err = fmt.Scanf("%d\n", &numero)
	fmt.Println(n, err, numero)
	if numero > 10 {
		fmt.Println(numero, "> 10")
	} else {
		fmt.Println("no es cierto que", numero, "> 10")
	}

	/*
		Operadores logicos && ||
	*/
}
