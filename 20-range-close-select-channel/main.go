package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	// Reserva espacio para dos mensajes
	c := make(chan string, 3)
	// Introduce 2 mensajes
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	// Cuantas datos tiene el channel, cuanta capacidad tiene
	fmt.Println(len(c), cap(c))

	// Cerrar canal
	// No recibe mas datos aunque tenga mas capacidad
	// Ideal cerrar despues de usarlo
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)
	go message("Email 1", email1)
	go message("Email 2", email2)

	// No sabemos cual canal responde primero
	// Usamos select
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email1 recibido", m1)
		case m1 := <-email2:
			fmt.Println("Email2 recibido", m1)
		}
	}
}
