package main

import (
	"fmt"
)

// definindo uma interface de tipo dinamico
type dinamico interface{}

type curso struct {
	nome string
}

func main() {
	var coisa dinamico
	var curso dinamico = curso{"go"}
	fmt.Println(curso)
	fmt.Println(coisa)
	coisa = 3
	fmt.Println(coisa)
	coisa = "string"
	fmt.Println(coisa)
	coisa = true
	fmt.Println(coisa)

}

/*
Go é fortemente tipado mas existem exceções
se eu criar uma interface dinamica(sem tipo),
posso declarar variaveis como se a linguagem não fosse tipada
*/
