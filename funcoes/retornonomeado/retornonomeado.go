package main

import "fmt"

//retornos nomeados
func trocar(p1, p2 int) (segundo, primeiro int) {
	// note que definimos nomes para os retornos
	segundo = p2
	primeiro = p1
	return // retorno limpo pois irá retornar os nomes previamentes definidos como retorno
}

func main() {
	x, y := trocar(1, 2)
	fmt.Println(x, y)

}
