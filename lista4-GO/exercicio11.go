package main

import "fmt"

func main() {
	var vetor [100]float64

	fmt.Println("Digite 100 valores numéricos:")
	for i := 0; i < 100; i++ {
		fmt.Printf("Valor %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	soma := 0.0
	for i := 0; i < 50; i++ {
		diff := vetor[i] - vetor[99-i]
		soma += diff * diff
	}

	fmt.Printf("\nSomatório S = %.2f\n", soma)
}
