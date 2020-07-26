package main

import "fmt"

//declarando uma estrutura
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//criando metodo p/ estrutura com funcao receiver
func (p produto) precoComDesconto() float64 {
	// note que: func (apelido.struct nome.struct) nome.do.metodo() tipo.de.retorno()
	return p.preco * (1 - p.desconto)
}

func main() {
	// instanciando estrutura
	var prod1 produto
	prod1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(prod1)
	//chamando metodo da struct
	fmt.Println(prod1.precoComDesconto())

	// outra maneira de instanciar uma struct

	prod2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(prod2.precoComDesconto())
}
