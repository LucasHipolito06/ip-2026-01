package main

import "fmt"

func main() {
	var vetor [10]int

	fmt.Println("Digite 10 números inteiros em ordem crescente:")
	for i := 0; i < 10; i++ {
		for {
			fmt.Printf("Posição %d: ", i)
			fmt.Scan(&vetor[i])
			if i == 0 || vetor[i] >= vetor[i-1] {
				break
			}
			fmt.Printf("Erro! O valor deve ser >= %d. Digite novamente.\n", vetor[i-1])
		}
	}

	fmt.Println("\nVetor ordenado crescente:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", vetor[i])
	}
	fmt.Println()
}
