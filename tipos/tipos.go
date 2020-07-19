package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// tipos numericos inteiros, funcao reflect.TypeOf retorna o tipo do dado
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (apenas positivos) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal  int6 int16 int32 int64
	i1 := math.MaxInt64 //pegar o maior valor inteiro 64bt
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("O tipo de i1 é ", reflect.TypeOf(i1))
	//representa um mapeamento da tabela Unicode(int32), ou seja o valor da string a na tabela unicode
	var i2 rune = 'a'
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é ", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá mundo !
	exemplo
	de
	quebra 
	linha com crase`
	fmt.Println(s2)

}
