package main

import "runtime/debug"

//pilhas de execução : first in last out
func f3() {
	//imprime a pilha de execução
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
