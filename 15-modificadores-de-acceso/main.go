package main

import (
	"fmt"
	"github.com/cdgn-coding/go-basico/15-modificadores-de-acceso/cars"
)

func main() {
	myCar := cars.Car{}
	myCar.Brand = "Ford"
	// myCar.code is not accesible from outside the package!
	fmt.Println(myCar)
	myCar.Move()
}
