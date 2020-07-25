package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int //slice vazio

	slice1 = append(slice1, 4, 5, 6) // recebendo append dele mesmo(vazio) + valores
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) //criando novo slice a partir do slice maker com 2 indices vazios
	copy(slice2, slice1)     // usando copy(slice.receptor,slice.a.ser.copiado) *obrigatoriamente tem que ser um slice ou string
	fmt.Println(slice2, slice1)
}
