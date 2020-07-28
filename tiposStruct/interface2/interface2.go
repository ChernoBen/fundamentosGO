package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	//turboLigado instanciado false
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()
	fmt.Println(carro1)
	// instanciando struct em var tipo sportivo
	var carro2 esportivo = &ferrari{"F40", false, 0}
	//note que terá um erro caso a variavel.tipo.esportivo nao receba um endereco de memoria por causa do receiver do metodo ligarTurbo que recebe um ponteiro
	fmt.Println(carro2)

}
