package main

import (
	"fmt"
	"math"
)

func main() {
	var numeros [15]int
	var raizes [15]float64

	fmt.Println("Digite 15 números inteiros:")
	for i := 0; i < 15; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])

		if numeros[i] < 0 {
			raizes[i] = -1
		} else {
			raizes[i] = math.Sqrt(float64(numeros[i]))
		}
	}

	fmt.Println("\nVetor com as raízes quadradas:")
	for i := 0; i < 15; i++ {
		fmt.Printf("Posição %d: %.4f\n", i, raizes[i])
	}
}
