package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Ciclo for... i=%d\n", i)
	}

	var counter = 0
	for counter < 10 {
		fmt.Printf("Ciclo while... counter=%d\n", counter)
		counter++
	}

	counter = 0
	for {
		if counter == 10 {
			break
		}
		fmt.Println("For forever... counter=", counter)
		counter++
	}
}
