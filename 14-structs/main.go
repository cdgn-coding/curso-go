package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	// Instancia asignando valores
	myCar := car{brand: " Ford", year: 2020}
	fmt.Println(myCar)

	// Instanciacion Vacio. Se inicializa con zero values
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
