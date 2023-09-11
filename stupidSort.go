package main

import (
	"math/rand"
	"time"
)

func verificar(arreglo []int) bool {
	for i := 0; i < len(arreglo)-1; i++ {
		if arreglo[i] > arreglo[i+1] {
			return false
		}
	}
	return true
}

// StupidSort
func Ordenador(arreglo []int) {
	rand.Seed(time.Now().UnixNano())
	for !verificar(arreglo) {
		posA := rand.Intn(len(arreglo))
		posB := rand.Intn(len(arreglo))

		aux := arreglo[posA]
		arreglo[posA] = arreglo[posB]
		arreglo[posB] = aux
	}
}
