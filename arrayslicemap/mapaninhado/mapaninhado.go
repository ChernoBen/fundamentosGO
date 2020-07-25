package main

import "fmt"

func main() {
	/*
		Declarando map com chave tipo string e valor= map com chave do tipo string e valor float64
	*/
	funcionarios := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15464.75,
			"Guga Pereira":   8456.45,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcionarios, "P")

	//imprimindo map aninhado usando um for aninhado
	for chave, valor := range funcionarios {
		fmt.Println(chave)
		for cha2, valor2 := range valor {
			fmt.Println(cha2, valor2)
		}

	}
}
