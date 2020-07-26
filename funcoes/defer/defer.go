package main

import "fmt"

//defer faz com que a linha especifica seja executata no ultimo momento possivel dentro e um dado metodo
func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim")
	if numero > 5000 {
		fmt.Println("Valor muito alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
