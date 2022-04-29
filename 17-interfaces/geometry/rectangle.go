package geometry

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Base, Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Base * rectangle.Height
}

func (rectangle Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Base: %.2f, Height: %.2f}", rectangle.Base, rectangle.Height)
}

func (rectangle Rectangle) Diagonal() float64 {
	return math.Sqrt(math.Pow(rectangle.Base, 2) + math.Pow(rectangle.Height, 2))
}
