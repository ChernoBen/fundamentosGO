package main

import (
	"fmt"
	m "math"
	/*posso referenciar uma biblioteca da maneira acima*/)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	//forma reduzida e declarar variaveis em Go
	area := PI * m.Pow(raio, 2)
	fmt.Println("Area da circunferencia é :", area)

	//declarando multiplas constantes
	const (
		a = 1
		b = 2
	)

	// declarando multiplas variaveis
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//declarando variaveis em uma unica linha
	var e, f bool = true, false
	fmt.Println(e, f)

	//declarando usando sintaxe reduzida( := é usado para declarar e inicializar a variavel)
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
