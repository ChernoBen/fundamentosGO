package main

import "fmt"

func main() {
	//arrays são estruturas homogeneas com mesmo tipo e estátiva não mutavel
	var notas [3]float64 //array inicializado com 3 indices com valores zerados
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	// calculando media a partir dos valores no array

	total := 0.0
	//também posso fazer assim : var total float64
	for i := 0; i < len(notas); i++ {
		total += notas[i]

	}
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
