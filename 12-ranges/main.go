package main

import (
	"fmt"
	"strings"
)

func palindromo(text string) bool {
	var max = len(text) - 1
	for i := 0; i < max/2; i++ {
		if text[i] != text[max-i] {
			return false
		}
	}
	return true
}

func main() {
	slice := []string{"hola", "que", "hace", "ama", "amma", "Amma"}
	for _, valor := range slice {
		fmt.Printf("%s es palindromo? %t\n", valor, palindromo(strings.ToLower(valor)))
	}
}
