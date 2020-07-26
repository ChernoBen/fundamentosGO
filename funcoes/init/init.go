package main

import "fmt"

//é uma convecao de go inicializar a funcao init primeiro
func init() {
	fmt.Println("inicializando...")
}

func main() {
	fmt.Println("main ...")
}
