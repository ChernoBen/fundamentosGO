package main

import (
	"fmt"
	"reflect"
)

// funcao que recebe multiplos valores dentro de um slice do tipo float
func media(numeros ...float64) float64 {
	var total float64
	fmt.Println(reflect.TypeOf(numeros))
	for _, valor := range numeros {
		total += valor
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
