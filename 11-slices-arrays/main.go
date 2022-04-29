package main

import "fmt"

func main() {
	// fixed size
	// memory already allocated
	// immutable
	var array [4]int
	// initialized with zero
	array[0] = 1
	fmt.Println(array, len(array), cap(array))

	// variable size
	// already allocated
	// mutable
	var slice = []int{1, 2, 3}
	fmt.Println(slice, len(slice), cap(slice))

	// allocating memory with slice
	allocatedSlice := make([]int, 3)
	// initialized with zero
	allocatedSlice[0] = 1
	fmt.Println(allocatedSlice, len(allocatedSlice), cap(allocatedSlice))

	// Append, agregar elementos
	slice = append(slice, 7)
	// Agregar una lista con el operador unpack
	slice = append(slice, []int{1, 2, 3, 4, 5}...)

	// slicing
	// primer elemento
	fmt.Println(slice[0])
	// desde i=0 con i<3
	fmt.Println(slice[:3])
	// desde i=2 con i<3
	fmt.Println(slice[2:3])
	// desde i=1 con i<len(slice)
	fmt.Println(slice[1:])
}
