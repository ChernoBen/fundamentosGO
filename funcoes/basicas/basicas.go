package main

import "fmt"

// ex de funcao sem paramentro e sem retorno
func f1() {
	fmt.Println("F1")
}

// ex funcao com parametro sem retorno
func f2(p1 string, p2 string) {
	//declarar parametro + tipo do parametro
	fmt.Printf("F2: %s %s \n", p1, p2)
}

// ex funcao sem parametro que retorna um valor , nesse caso tipo string
func f3() string {
	return "F3"
}

//ex funcao que recebe parametro e retorna valor formatado
func f4(p1, p2 string) string {
	//metodo sprint não imprimi mas retorna uma string formatada
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// ex funcao sem parametro que retorna 2 parametros
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param 1", "Param 2")
	fmt.Println(f3())
	fmt.Println(f4("Param 1", "Param 2"))
	//f5() retorna 2 parametros entao:
	r1, r2 := f5()
	fmt.Println(r1, r2)

}
