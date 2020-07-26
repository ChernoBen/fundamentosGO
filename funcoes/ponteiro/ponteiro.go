package main

import "fmt"

// ponteiros em chamadas de funcoes
func inc(n int) {
	n++
}

//lembrar que ponteiro é representado por um *
func inc2(n *int) {
	// nesse caso o * é usado para desreferenciar e obter o valor que ele referencia
	*n++
}
func main() {
	n := 1

	inc(n) // passagem por valor
	fmt.Println(n)

	//& usado para obter o endereco da variavel oque vai causar mudancas externar a funcao
	inc2(&n)
	fmt.Println(n)
}

//no primeiro caso variavel n nao sofreu alteracao
//no segundo sim pois o endereco de memoria foi passado
