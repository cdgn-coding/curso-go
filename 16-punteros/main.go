package main

import "fmt"

type pc struct {
	ram   int
	brand string
}

func (pc pc) String() string {
	return fmt.Sprintf("marca %s con ram %d gb", pc.brand, pc.ram)
}

func main() {
	a := 50
	// Direccion de a
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = *b + 50
	fmt.Println(*b)

	myPc := pc{ram: 16, brand: "msi"}
	fmt.Println(myPc)
}
