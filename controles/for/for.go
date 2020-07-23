package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { //Go nao possui while
		fmt.Println(i)
		i++
	}
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ,", i)
	}
	for {
		//laço infinito
		fmt.Println("'FOR' ever ...")
		time.Sleep(time.Second * 2)
	}
}
