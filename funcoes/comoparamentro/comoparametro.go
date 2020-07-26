package main

import "fmt"

// funcao como parametro de funcao
func multiplicacao(a, b int) int {
	return a * b
}

//funcao que recebe(func(tipo.int,tipo.int)retorna.tipo.int, e recebe 2 parametros(p1,p2 do tipo int) e retorna valor inteiro)
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 2, 4)
	fmt.Println(resultado)
}
