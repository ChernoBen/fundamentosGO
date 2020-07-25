package main

import "fmt"

func main() {
	//declarando e inicializando um map
	funcionarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	//adionando novo funcionario
	funcionarios["Rafael Filho"] = 1350.0
	for chave, valor := range funcionarios {
		fmt.Println(chave, valor)
	}
}
