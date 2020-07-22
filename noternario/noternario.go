package main

import "fmt"

// alternativa ao operador ternario

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "aprovador"
	}
	return "resultado"

}
func main() {
	fmt.Println(obterResultado(6.2))
}
