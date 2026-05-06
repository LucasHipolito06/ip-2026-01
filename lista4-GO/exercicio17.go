package main

import (
	"fmt"
	"math"
)

func main() {
	var vetor [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor[i])
	}

	fmt.Println("\nNúmeros primos e suas posições:")
	encontrou := false
	for i := 0; i < 10; i++ {
		if ehPrimo(vetor[i]) {
			fmt.Printf("Valor: %d - Posição: %d\n", vetor[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhum número primo encontrado.")
	}
}

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	limite := int(math.Sqrt(float64(n)))
	for i := 3; i <= limite; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
