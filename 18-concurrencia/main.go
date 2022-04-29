package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("hello")
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Adios")
	}()

	wg.Wait()
}
