package main

import "fmt"

func main() {
	var vetor [30]int
	var resultado [30]int

	fmt.Println("Digite 30 números inteiros:")
	for i := 0; i < 30; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			resultado[i] = vetor[i] * 2
		} else {
			resultado[i] = vetor[i] * 3
		}
	}

	fmt.Println("\nVetor original:", vetor)
	fmt.Println("Vetor resultante:", resultado)
}
