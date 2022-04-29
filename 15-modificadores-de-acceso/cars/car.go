package cars

import "fmt"

type Car struct {
	Brand string
	Year  int
	code  string
}

func (c Car) Move() {
	fmt.Println("I am moving!")
}
