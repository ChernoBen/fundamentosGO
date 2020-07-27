package main

import "fmt"

// struct p/ guardar info de produtos
type item struct {
	produtoID int
	qtd       int
	preco     float64
}

// struct para organizar pedidos
type pedido struct {
	userID int
	//slice de itens da struct item
	itens []item
}

// metodo p/ processar valores em 2 estuturas
func (p pedido) valorTotal() float64 {
	var total float64
	//p/ cada valor dentro do slice pedido.itens
	for _, val := range p.itens {
		//valor.preco X valor.quatidade
		total += val.preco * float64(val.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			//item{produtoID: 1, qutd: 2, preco: 3.4}
			item{1, 2, 3.4},
			item{4, 5, 6.7},
			item{8, 9, 10.11},
		},
	}
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
	for _, vl := range pedido.itens {
		fmt.Println("\n", vl)
	}
}
