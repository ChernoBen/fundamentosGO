package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	//bloco de inicialização antes de executar o if + condição de entrada
	if i := numeroAleatorio(); i > 5 {
		// se o numero atribuido em 'i' for maior que 5 faça:
		fmt.Println("ganhou")

	} else {
		fmt.Println("Perdeu")
	}
}
