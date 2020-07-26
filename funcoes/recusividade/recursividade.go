package main

import "fmt"

// funcao recursiva que retorna valor inteiro e um erro
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		//o segredo da recursividade ta bem aqui na variavel recebendo a chamada da propria funcão
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
