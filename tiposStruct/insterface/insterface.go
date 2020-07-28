package main

import "fmt"

// criando um tipo imprivivel
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobremone string
}

type produto struct {
	nome  string
	preco float64
}

// metodo p/ struct pessoa
func (p pessoa) toString() string {
	return p.nome + " " + p.sobremone
}

// metodo para retornar itens da struct em forma de string
func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	//Instanciando struct em variavel do tipo imprimivel
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	// ou apenas
	fmt.Println(coisa)
	// ou tb
	imprimir(coisa)
}
