package main

import (
	"fmt"
)

func main() {

	listaDes := []int{1, 5, 2, 3, 7, 0, 4, 8, 15}

	fmt.Println("Lista desordenada", listaDes)
	Ordenador(listaDes)
	fmt.Println("Lista ordenada", listaDes)

}
