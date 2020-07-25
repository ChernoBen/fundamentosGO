package main

import "fmt"

func main() {
	// criando um array dinâmico
	numeros := [...]int{1, 2, 3, 4, 5} // [...] espaco será definido com base nos valores que eu adicionar
	//note que dessa maneira o laço coleta index a value do array, pode ser muito util para percorrer maps(ou dicionarios)
	for indice, valor := range numeros {
		fmt.Printf("%d) %d \n", indice, valor)
	}

	// também posso ignorar o indice
	for _, value := range numeros {
		fmt.Println(value)
	}
}
