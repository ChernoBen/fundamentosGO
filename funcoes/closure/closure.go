package main

import "fmt"

//funcao sem parametros que retorna uma funcao
func clousure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}
func main() {
	// a variavel imprimeX ira imprimir 10 pois o valor está definido na declaracao da função
	x := 20
	fmt.Println(x)
	imprimeX := clousure()
	imprimeX()
}
