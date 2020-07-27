package main

import "fmt"

type carro struct {
	nome             string
	velocidadeAtrual int
}

type ferrari struct {
	carro       //campo anonimo, todas as caracteristicas de carro agora esta disponivel em ferrari/composicao
	turboLigado bool
}

func main() {
	// instanciando struct pseudo-filha
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtrual = 60
	f.turboLigado = true
	fmt.Printf(" %s: ligada \n Turbo ligado: %v \n Velocidade atual: %d", f.nome, f.turboLigado, f.velocidadeAtrual)

}
