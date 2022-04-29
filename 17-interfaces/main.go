package main

import (
	"fmt"
	"github.com/cdgn-coding/go-basico/17-interfaces/geometry"
)

func main() {
	rectangle := geometry.Rectangle{Base: 5, Height: 10}
	var figures = []geometry.Figure2d{
		geometry.Circle{Radius: 5},
		rectangle,
	}

	for _, figure := range figures {
		fmt.Println(figure, "Area=", figure.Area())
	}

	fmt.Println(rectangle, "Diagonal=", rectangle.Diagonal())

	// Interface sin restricciones
	interfaces := []interface{}{"Hola", 123, 4.99}
	fmt.Println(interfaces...)
}
