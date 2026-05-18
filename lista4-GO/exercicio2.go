package main

import "fmt"

func main() {
	var vetor1 [10]int
	var vetor2 [5]int

	fmt.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("Digite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor2[i])
	}

	somaVetor2 := 0
	for i := 0; i < 5; i++ {
		somaVetor2 += vetor2[i]
	}

	var resultPares []int
	var resultImpares []int

	for i := 0; i < 10; i++ {
		if vetor1[i]%2 == 0 {
			resultPares = append(resultPares, vetor1[i]+somaVetor2)
		} else {
			resultImpares = append(resultImpares, vetor1[i]+somaVetor2)
		}
	}

	fmt.Println("\nVetor resultante (pares):", resultPares)
	fmt.Println("Vetor resultante (ímpares):", resultImpares)
}
