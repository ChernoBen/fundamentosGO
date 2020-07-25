package main

import "fmt"

func main() {
	//criando slice com função make
	s := make([]int, 10) //make([slice]tipo.Int,10-indices)
	s[9] = 12
	fmt.Println(s)

	//reatribuindo valor a s
	s = make([]int, 10, 20)        //make([slice]tipo.int,10-indices,20-indices no arrey interno)
	fmt.Println(s, len(s), cap(s)) //slice.S, tamanho.Slice.S, capacidade.do.array.de.S

	//note que o slice guarda o valor 1 em um novo indice
	//caso a quantidade seja maior que a quantidade de inideces ele dobrará o tamno do array interno
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1)
	fmt.Println(s, len(s), cap(s))
}
