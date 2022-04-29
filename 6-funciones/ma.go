package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hola %s\n", name)
}

func area(b, h float64) float64 {
	return b * h
}

func main() {
	greet("Carlos")
	greet("Johana")
	greet("Carmen")
	fmt.Printf("Area de rectangulo 3 x 4 = %.2f\n", area(3, 2))
	var b, h float64
	var n int
	var err error
	fmt.Print("Ingrese base de triangulo: ")
	n, err = fmt.Scanf("%f\n", &b)
	fmt.Println(n, err, b)

	fmt.Print("Ingrese altura de triangulo: ")
	_, _ = fmt.Scanf("%f\n", &h)
	fmt.Println(n, err, h)

	fmt.Printf("Area de rectangulo %.2f x %.2f = %.2f", b, h, area(b, h))

}
