package main

import "fmt"

//compor uma nova interface a partir de 2 outras interfaces
type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerbaliza()
}

//composicao de interfaces
type esportivoluxuoso interface {
	luxuoso
	esportivo
}

type bmw7 struct {
}

func (b bmw7) ligarTurbo() {
	fmt.Println("TurboLigado...")
}

func (b bmw7) fazerbaliza() {
	fmt.Println("fazendo baliza...")
}

func main() {
	var b esportivoluxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerbaliza()
}

// obs sobre composicao e interfaces:
/*
Pode ser uma otima opção para processamento de dados.
Defininto tipos de interfaces, posso usar como filtro pra separar determinados dados .
*/
