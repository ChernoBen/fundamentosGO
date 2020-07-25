package main

import "fmt"

func main() {
	// criando e inicializando um map(dicionario) com a função make
	aprovados := make(map[string]int)

	aprovados["Maria"] = 123
	aprovados["Pedro"] = 456
	aprovados["Ana"] = 789
	fmt.Println(aprovados)

	//usando for para imprimir cada item do map(dicionario)
	for chave, valor := range aprovados {
		fmt.Printf(" %s (CPF: %d ) \n", chave, valor)
	}
	//removendo elemento de um map
	delete(aprovados, "Maria") //apresentar map e chave do map
	fmt.Println(aprovados["Maria"])
}
