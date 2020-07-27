package main

import "fmt"

// tipo identificador e tipo base
type nota float64

//gerando funcao que recebe como receiver meu.tipo + funcao verificadora que retorna um bool
func (n nota) entre(inicio, fim float64) bool {
	// retornar se :
	return float64(n) >= inicio && float64(n) <= fim

}

// funcao que recebe um parametro do tipo nota
func notaPraConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaPraConceito(9.8))
	fmt.Println(notaPraConceito(6.9))
}
