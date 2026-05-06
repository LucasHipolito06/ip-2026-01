package main

import "fmt"

func main() {
	var vetor [100]int

	for i := 0; i < 100; i++ {
		vetor[i] = 100 - i
	}

	fmt.Println("Vetor com números de 100 a 1 (ordem decrescente):")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", vetor[i])
	}
	fmt.Println()
}
