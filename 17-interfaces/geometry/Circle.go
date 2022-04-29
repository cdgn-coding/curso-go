package geometry

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", circle.Radius)
}
