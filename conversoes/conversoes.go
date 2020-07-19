package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("valor da nota final é", notaFinal)

	//note que ele não imprimirá 97 como string, mas sim o valor correspondente na tabela asc
	fmt.Println("teste" + string(97))

	// conveter inteiro p/ string
	fmt.Println("Teste" + strconv.Itoa(123))
	// metodo inverso
	num, _ := strconv.Atoi("123")
	/*
		esse metodo retorna 2 valores e _ é usado para ignorar o valor erro retornado
	*/
	fmt.Println(num - 122)
	//conversao p/ boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("verdadeiro")
	}
}
