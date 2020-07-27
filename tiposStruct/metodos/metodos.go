package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

// metodo p/ obter nome completo a partir da struct
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// alterar nome completo a partir de um ponteiro que aponta p/ valores dentro da struct
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	// instanciando struct
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())
	p1.setNomeCompleto("Benjamim Francisco")
	fmt.Println(p1.getNomeCompleto())
}
