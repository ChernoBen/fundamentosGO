package main

import "fmt"

func main() {
	/*
		Declarar variaveis zeradas, tenho que passar o tipo, pq não atribuí nenhum valor
	*/
	var a int     // por padrão é 0
	var b float64 // por padrão é 0
	var c bool    // por padrão false
	var d string  //vazio
	var e *int    // <nil> que é null para ponteiro

	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)

}
