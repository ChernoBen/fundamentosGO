package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	//note a diferença quando:
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))
	//slice é um pedaço referenciado de um array
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] //defindo um slice do indice 1 ao 2, ignorando 3
	fmt.Println(a2, s2)
	s3 := a2[:2] // novo slice aponta pro inicio do array até indice 2 mas ignorando 2
	fmt.Println(s3)
	s4 := s2[:1] //slice de um slice
	fmt.Println(s4)
}
