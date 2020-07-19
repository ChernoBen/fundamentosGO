package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print(" linha.")
	// imprimir com quebra de linha
	fmt.Println(" Nova")
	fmt.Println("Linha.")

	x := 3.141516
	fmt.Println("O valor x  é ", x)

	// retorna uma string
	xs := fmt.Sprint(x)

	fmt.Println("o valor de x versão string é " + xs)

	//print formatado
	fmt.Printf("o valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
