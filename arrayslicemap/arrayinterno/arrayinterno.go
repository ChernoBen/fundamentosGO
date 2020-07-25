package main

import "fmt"

func main() {

	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	/*se os 2 slices estão apontando de fato p/ o mesmo array,
	então o valor 7 estrá presente nos mesmo indices dos 2 slices */

	s1[0] = 7
	fmt.Println(s1, s2)
	//note que de fato os 2 slices apontam para o mesmo array

}
