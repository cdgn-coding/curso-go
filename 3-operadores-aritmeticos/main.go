package main

import "fmt"

func main() {
	x := 10
	y := 50

	fmt.Println("x + y", x+y)
	fmt.Println("x * y", x*y)
	fmt.Println("x - y", x-y)
	fmt.Println("y / x (entera)", y/x)
	fmt.Println("y % x", y%x)
	x++
	fmt.Println("x++", x)
	x--
	fmt.Println("x--", x)
}
