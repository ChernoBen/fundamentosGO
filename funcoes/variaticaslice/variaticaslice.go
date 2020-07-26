package main

import "fmt"

// funcao variatica com slice como parametro
func imprimiralunos(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for chave, valor := range aprovados {
		fmt.Printf("%d) %s \n", chave+1, valor)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	//passar indices do slice como parametos da funcao variatica (slice...)
	imprimiralunos(aprovados...)
}
