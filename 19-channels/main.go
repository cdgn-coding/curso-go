package main

import (
	"fmt"
	"time"
)

func say(text string, c chan<- string) {
	fmt.Println(text)
	c <- "Resultado"
}

func main() {
	// chan es tipo channel
	// len es cuantos datos simultaneos maneja
	// si no lo ponemos, se asume 1 solo dato
	c := make(chan string, 1)

	fmt.Println("Hello")
	// Se espera la ejecucion si se lee el valor del channel
	go say("Bye", c)
	fmt.Println(<-c)
	time.Now()
}
