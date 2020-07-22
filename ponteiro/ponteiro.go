package main

import "fmt"

func main() {
	i := 1

	// Go nao possui aritimética de ponteiros ex: p++ ou p +=2 etc...
	var p *int = nil // nil == nulo
	p = &i           // atribuindo endereco da variavel i e atribuindo p/ ponteiro p
	*p++             // desreferenciando o ponteiro para obter o valor
	i++

	fmt.Println("endereco :", p, "valor dentro do endereco :", *p, "\n valor dentro do endereco :", i, "endereco :", &i)
}
