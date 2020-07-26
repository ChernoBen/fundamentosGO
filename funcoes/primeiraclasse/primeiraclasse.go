package main

import "fmt"

// funoes anonimas GO/ variavel que recebe uma funcao anonima
var soma = func(a, b int) int {
	return a + b
}

func main() {
	//executando funcao anonima dentro da variavel
	fmt.Println(soma(5, 5))

	// declarando o mesmo tipo de funcao dentro da funcao principal
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(10, 5))
}
