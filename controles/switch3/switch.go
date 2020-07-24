package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string { // interface quer dizer que aceita tipos variados
	switch i.(type) { // se o tipo for :
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case func():
		return "Função"
	default:
		return "Não sei"

	}
}
func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
